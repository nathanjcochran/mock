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
		return Interface{}, fmt.Errorf("erroring parsing directory: %s", err)
	}

	for pkgPath, pkgAST := range pkgASTs {
		var (
			files       = []*ast.File{}
			fileImports = map[token.Pos][]Import{}
		)
		for _, file := range pkgAST.Files {
			files = append(files, file)

			// Keep track of each file's imports, and
			// their names, if they were renamed:
			var imprts []Import
			for _, imp := range file.Imports {
				imprt := Import{
					Path: strings.Trim(imp.Path.Value, "\""),
				}
				if imp.Name != nil {
					imprt.Name = imp.Name.Name
				}
				imprts = append(imprts, imprt)
			}
			fileImports[file.Pos()] = imprts
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
			return Interface{}, fmt.Errorf("%s is not a named type", ifaceName)
		}
		ifaceType, ok := ifaceObj.Type().Underlying().(*types.Interface)
		if !ok {
			return Interface{}, fmt.Errorf("%s is not an interface", ifaceName)
		}

		// Get the AST nodes enclosing the interface type name object:
		fileName := fset.File(ifaceObj.Pos()).Name()
		ifaceFile, exists := pkgAST.Files[fileName]
		if !exists {
			return Interface{}, fmt.Errorf("Could not find file: %s", fileName)
		}
		path, _ := astutil.PathEnclosingInterval(ifaceFile, ifaceObj.Pos(), ifaceObj.Pos())

		// Find the interface definition/declaration in the AST:
		var ifaceDecl *ast.GenDecl
		for _, node := range path {
			if ifaceDecl, ok = node.(*ast.GenDecl); ok {
				break
			}
		}
		if ifaceDecl == nil {
			return Interface{}, fmt.Errorf("Could not find interface declaration in AST")
		}

		// If there were type errors, make sure none of them were relevant
		// to the interface definition, or return the first that was:
		for _, err := range typeErrs {
			typErr, ok := err.(types.Error)
			if !ok {
				return Interface{}, err
			}

			// Return the first hard error relevant to the interface:
			if !typErr.Soft &&
				typErr.Pos >= ifaceDecl.Pos() &&
				typErr.Pos <= ifaceDecl.End() {
				return Interface{}, typErr
			}
		}


		// Get the file's imports:
		fileImprts := fileImports[fset.File(ifaceObj.Pos()).Pos(0)]

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
					Type: types.TypeString(paramObj.Type(), Qualify(pkg, fileImprts, &iface.Imports)),
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
					Type: types.TypeString(resultObj.Type(), Qualify(pkg, fileImprts, &iface.Imports)),
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

func Qualify(pkg *types.Package, fileImprts []Import, usedImprts *[]Import) types.Qualifier {
	return func(other *types.Package) string {
		// If the type is from this package, don't qualify it:
		if pkg == other {
			return ""
		}

		// Search for the import statement for the package
		// that the type is from:
		for _, imprt := range fileImprts {
			if other.Path() == imprt.Path {
				// If the package was only imported for its
				// side-effects, skip over it:
				if imprt.Name == "_" {
					continue
				}

				// Keep track of the file imports that have actually
				// been used in this interface definition (de-duped):
				var found bool
				for _, usedImprt := range *usedImprts {
					if imprt.Path == usedImprt.Path {
						found = true
						break
					}
				}
				if !found {
					*usedImprts = append(*usedImprts, imprt)
				}

				// If the package was brought into this package
				// in an unqualified manner, don't qualify it:
				if imprt.Name == "." {
					return ""
				}

				// If the package was renamed in the import
				// statement, return it's name:
				if imprt.Name != "" {
					return imprt.Name
				}

				// Othewise, the package was not renamed,
				// so break out and return the package name:
				break
			}
		}
		return other.Name()
	}
}
