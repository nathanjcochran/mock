package iface

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"sort"
	"strings"
)

func GetInterface(dir, ifaceName string) (Interface, error) {
	fset := token.NewFileSet()
	pkgASTs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return Interface{}, fmt.Errorf("erroring parsing .go files in directory: %s", err)
	}

	for pkgPath, pkgAST := range pkgASTs {
		var (
			files    = []*ast.File{}
			fileImps = map[token.Pos][]Import{}
		)
		for _, file := range pkgAST.Files {
			files = append(files, file)

			// Keep track of each file's imports, and
			// their names, if they were renamed:
			var imps []Import
			for _, fileImp := range file.Imports {
				imp := Import{
					Path: strings.Trim(fileImp.Path.Value, "\""),
				}
				if fileImp.Name != nil {
					imp.Name = fileImp.Name.Name
				}
				imps = append(imps, imp)
			}
			fileImps[file.Pos()] = imps
		}

		// Type-check the package:
		var typeErrs []error
		conf := types.Config{
			Error: func(err error) {
				typeErrs = append(typeErrs, err)
			},
			Importer: importer.For("source", nil),
		}
		pkg, _ := conf.Check(pkgPath, fset, files, nil)

		// Find the interface by name:
		ifaceObj := pkg.Scope().Lookup(ifaceName)
		if ifaceObj == nil {
			continue // Interface not found in this pkg, try next
		}

		// Validate that the object with that name
		// is indeed an interface:
		if _, ok := ifaceObj.(*types.TypeName); !ok {
			return Interface{}, fmt.Errorf("%s is not a named/defined type", ifaceName)
		}
		ifaceType, ok := ifaceObj.Type().Underlying().(*types.Interface)
		if !ok {
			return Interface{}, fmt.Errorf("%s is not an interface type", ifaceName)
		}

		// Make sure that none of the types involved in the
		// interface's definition were invalid/had errors:
		if !ValidateType(ifaceType) {
			return Interface{}, &TypeErrors{Errs: typeErrs}
		}

		// Get the file's imports:
		imps := fileImps[fset.File(ifaceObj.Pos()).Pos(0)]

		iface := Interface{
			Package: pkg.Name(),
			Name:    ifaceObj.Name(),
		}
		for i := 0; i < ifaceType.NumMethods(); i++ {
			methodObj := ifaceType.Method(i)
			sig, ok := methodObj.Type().(*types.Signature)
			if !ok {
				return Interface{}, fmt.Errorf("%s is not a method signature", methodObj.Name())
			}
			method := Method{
				Name: methodObj.Name(),
				pos:  methodObj.Pos(),
			}

			// Keep track of each of the method's parameters:
			paramsTuple := sig.Params()
			for j := 0; j < paramsTuple.Len(); j++ {
				paramObj := paramsTuple.At(j)
				param := Param{
					Name: paramObj.Name(),
					Type: types.TypeString(paramObj.Type(), Qualify(pkg, imps, &iface.Imports)),
				}
				method.Params = append(method.Params, param)
			}

			// Keep track of whether the last parameter is variadic:
			if len(method.Params) > 0 && sig.Variadic() {
				method.Params[len(method.Params)-1].Variadic = true
			}

			// Keep track of each of the method's results:
			resultsTuple := sig.Results()
			for j := 0; j < resultsTuple.Len(); j++ {
				resultObj := resultsTuple.At(j)
				result := Result{
					Name: resultObj.Name(),
					Type: types.TypeString(resultObj.Type(), Qualify(pkg, imps, &iface.Imports)),
				}
				method.Results = append(method.Results, result)
			}

			iface.Methods = append(iface.Methods, method)
		}
		sort.Sort(iface.Methods)

		return iface, nil
	}

	return Interface{}, fmt.Errorf("interface not found: %s", ifaceName)
}

func Qualify(pkg *types.Package, imps []Import, usedImps *[]Import) types.Qualifier {
	return func(other *types.Package) string {
		// If the type is from this package, don't qualify it:
		if pkg == other {
			return ""
		}

		// Search for the import statement for the package
		// that the type is from:
		for _, imp := range imps {
			if other.Path() == imp.Path {
				// If the package was only imported for its
				// side-effects, skip over it:
				if imp.Name == "_" {
					continue
				}

				// Keep track of the file imports that have actually
				// been used in this interface definition (de-duped):
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
				// in an unqualified manner, don't qualify it:
				if imp.Name == "." {
					return ""
				}

				// If the package was renamed in the import
				// statement, return it's name:
				if imp.Name != "" {
					return imp.Name
				}

				// Othewise, the package was not renamed,
				// so break out and return the package name:
				break
			}
		}
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
