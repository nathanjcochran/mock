package iface

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

func GetInterface(dir, ifaceName string) (Interface, error) {
	cfg := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		return Interface{}, errors.Wrap(err, "error loading package info")
	}

	if len(pkgs) < 1 {
		return Interface{}, errors.Wrap(err, "failed to find/load package info")
	} else if len(pkgs) > 1 {
		return Interface{}, errors.Wrap(err, "found more than one matching package")
	}
	pkg := pkgs[0]

	// Keep track of each file's imports, along with their name (if renamed)
	fileImps := map[token.Pos][]Import{}
	for _, fileAST := range pkg.Syntax {
		var imps []Import
		for _, fileImp := range fileAST.Imports {
			imp := Import{
				Path: strings.Trim(fileImp.Path.Value, "\""),
			}
			if fileImp.Name != nil {
				imp.Name = fileImp.Name.Name
			}
			imps = append(imps, imp)
		}
		fileImps[pkg.Fset.File(fileAST.Pos()).Pos(0)] = imps
	}

	// Find the interface by name
	ifaceObj := pkg.Types.Scope().Lookup(ifaceName)
	if ifaceObj == nil {
		return Interface{}, errors.Errorf("interface not found in package: %s", ifaceName)
	}

	// Validate that the object with that name
	// is indeed an interface
	if _, ok := ifaceObj.(*types.TypeName); !ok {
		return Interface{}, fmt.Errorf("%s is not a named/defined type", ifaceName)
	}
	ifaceType, ok := ifaceObj.Type().Underlying().(*types.Interface)
	if !ok {
		return Interface{}, fmt.Errorf("%s is not an interface type", ifaceName)
	}

	// Make sure that none of the types involved in the
	// interface's definition were invalid/had errors
	if !ValidateType(ifaceType) {
		return Interface{}, &TypeErrors{Errs: pkg.Errors}
	}

	// Get the file's imports
	imps := fileImps[pkg.Fset.File(ifaceObj.Pos()).Pos(0)]

	// Begin assembling information about the interface
	iface := Interface{
		Package: pkg.Name,
		Name:    ifaceObj.Name(),
	}

	// Iterate through each embedded interface's explicit methods
	for _, ifaceType := range explodeInterface(ifaceType) {
		for i := 0; i < ifaceType.NumExplicitMethods(); i++ {
			methodObj := ifaceType.ExplicitMethod(i)
			method := Method{
				Name:            methodObj.Name(),
				SourceInterface: ifaceType.String(),
				pos:             methodObj.Pos(),
			}

			sig, ok := methodObj.Type().(*types.Signature)
			if !ok {
				return Interface{}, fmt.Errorf("%s is not a method signature", methodObj.Name())
			}

			// Keep track of the names and types of the parameters
			paramsTuple := sig.Params()
			for j := 0; j < paramsTuple.Len(); j++ {
				paramObj := paramsTuple.At(j)
				param := Param{
					Name: paramObj.Name(),
					Type: types.TypeString(paramObj.Type(), Qualify(pkg.Types, imps, &iface.Imports)),
				}
				method.Params = append(method.Params, param)
			}

			// Mark whether the last parameter is variadic
			if len(method.Params) > 0 && sig.Variadic() {
				method.Params[len(method.Params)-1].Variadic = true
			}

			// Keep track of the names and types of the results
			resultsTuple := sig.Results()
			for j := 0; j < resultsTuple.Len(); j++ {
				resultObj := resultsTuple.At(j)
				result := Result{
					Name: resultObj.Name(),
					Type: types.TypeString(resultObj.Type(), Qualify(pkg.Types, imps, &iface.Imports)),
				}
				method.Results = append(method.Results, result)
			}

			iface.Methods = append(iface.Methods, method)
		}
	}

	// Preserve the original ordering of the methods
	sort.Sort(iface.Methods)

	return iface, nil
}

// explodeInterface traverses an interface type, returning the original
// interface along with all transitively embedded interfaces.
func explodeInterface(iface *types.Interface) []*types.Interface {
	var (
		result    []*types.Interface
		workQueue = []*types.Interface{iface}
		visited   = map[string]bool{}
	)
	for len(workQueue) > 0 {
		current := workQueue[0]
		workQueue = workQueue[1:]
		currentID := current.String()
		if !visited[currentID] {
			visited[currentID] = true
			result = append(result, current)
			for i := 0; i < current.NumEmbeddeds(); i++ {
				switch embeddedIface := current.EmbeddedType(i).(type) {
				case *types.Interface:
					workQueue = append(workQueue, embeddedIface)
				}
			}
		}
	}
	return result
}

func Qualify(pkg *types.Package, imps []Import, usedImps *[]Import) types.Qualifier {
	return func(other *types.Package) string {
		// If the type is from this package, don't qualify it
		if pkg == other {
			return ""
		}

		// Search for the import statement for the package
		// that the type is from
		for _, imp := range imps {
			if other.Path() == imp.Path {

				// If the package was only imported for its
				// side-effects, skip over it
				if imp.Name == "_" {
					continue
				}

				// Keep track of the file imports that have actually
				// been used in this interface definition (de-duped)
				var found bool
				for _, usedImprt := range *usedImps {
					if imp.Path == usedImprt.Path {
						found = true
						break
					}
				}
				if !found {
					*usedImps = append(*usedImps, imp)
				}

				// If the package was brought into this package
				// in an unqualified manner, don't qualify it
				if imp.Name == "." {
					return ""
				}

				// If the package was renamed in the import
				// statement, return it's name
				if imp.Name != "" {
					return imp.Name
				}

				// Othewise, the package was not renamed,
				// so break out and return the package name
				break
			}
		}

		// We were unable to find an import statement in the original
		// file containing the interface that corresponds to the type.
		// This can happen if, for example, the interface embeds another
		// type from a different file/package. Add a corresponding import
		// to the list of used imports.
		// TODO: Because this is a new import that's not coming from the
		// original file, it could cause naming conflicts
		*usedImps = append(*usedImps, Import{
			Path: other.Path(),
		})
		return other.Name()
	}
}

func ValidateType(typ types.Type) bool {
	return validateType(typ, &[]types.Type{})
}

func validateType(typ types.Type, visited *[]types.Type) bool {
	for _, t := range *visited {
		if t == typ {
			return true
		}
	}
	*visited = append(*visited, typ)

	switch t := typ.(type) {
	case nil:
		return true

	case *types.Basic:
		return t.Kind() != types.Invalid

	case *types.Array:
		return validateType(t.Elem(), visited)

	case *types.Slice:
		return validateType(t.Elem(), visited)

	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			if !validateType(t.Field(i).Type(), visited) {
				return false
			}
		}
		return true

	case *types.Pointer:
		return validateType(t.Elem(), visited)

	case *types.Tuple:
		for i := 0; i < t.Len(); i++ {
			if !validateType(t.At(i).Type(), visited) {
				return false
			}
		}
		return true

	case *types.Signature:
		return validateType(t.Params(), visited) &&
			validateType(t.Results(), visited)

	case *types.Interface:
		for i := 0; i < t.NumMethods(); i++ {
			if !validateType(t.Method(i).Type(), visited) {
				return false
			}
		}
		return true

	case *types.Map:
		return validateType(t.Elem(), visited) &&
			validateType(t.Key(), visited)

	case *types.Chan:
		return validateType(t.Elem(), visited)

	case *types.Named:
		return validateType(t.Underlying(), visited)

	default:
		// log.Printf("Unknown types.Type: %v", t)
		return true
	}
}
