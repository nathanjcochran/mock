package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	dir := flag.String("dir", "", "Directory to search for interface in")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough args")
	}

	if *dir == "" {
		// Default to current working directory:
		var err error
		*dir, err = os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory: %s", err)
		}
	}

	intfName := args[0]
	intf, err := GetInterface(*dir, intfName)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("default").Parse(defaultTmpl)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	if err := tmpl.Execute(os.Stdout, &intf); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}
}

type Interface struct {
	PackageName   string
	InterfaceName string
	Methods       []Method
	Imports       []Import
}

type Method struct {
	Name    string
	Params  []Param
	Results []Result
}

type Param struct {
	Name     string
	Type     string
	Variadic bool
}

type Result struct {
	Name string
	Type string
}

type Import struct {
	Name string
	Path string
}

func (i *Interface) ImportsString() string {
	var strs []string
	for _, imprt := range i.Imports {
		var str string
		if imprt.Name != "" {
			str = fmt.Sprintf("\t%s \"%s\"", imprt.Name, imprt.Path)
		} else {
			str = fmt.Sprintf("\t\"%s\"", imprt.Path)
		}
		strs = append(strs, str)
	}
	return strings.Join(strs, "\n")
}

// TODO: Handle unnamed params, and blank-identifier params
func (m *Method) ParamsString() string {
	var strs []string
	for _, param := range m.Params {
		typ := param.Type
		if param.Variadic {
			typ = fmt.Sprintf("...%s", param.Type) // TODO: is type already a slice?
		}

		str := typ
		if param.Name != "" {
			str = fmt.Sprintf("%s %s", param.Name, typ)
		}
		strs = append(strs, str)
	}
	return strings.Join(strs, ", ")
}

func (m *Method) ResultsString() string {
	var strs []string
	for _, result := range m.Results {
		str := result.Type
		if result.Name != "" {
			str = fmt.Sprintf("%s %s", result.Name, result.Type)
		}
		strs = append(strs, str)
	}
	str := strings.Join(strs, ", ")
	if len(strs) < 2 {
		return str
	}
	return fmt.Sprintf("(%s)", str)
}

func GetInterface(dir, intfName string) (Interface, error) {
	fset := token.NewFileSet()
	pkgASTs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return Interface{}, fmt.Errorf("Erroring parsing file: %s", err)
	}

	conf := types.Config{Importer: importer.Default()}
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
		pkg, err := conf.Check(pkgPath, fset, files, nil)
		if err != nil {
			return Interface{}, fmt.Errorf("Type error: %s", err)
		}

		// Find the interface by name:
		intfObj := pkg.Scope().Lookup(intfName)
		if intfObj == nil {
			continue // Interface not found in this pkg, try next
		}

		// Validate that the object with that name
		// is indeed an interface:
		if _, ok := intfObj.(*types.TypeName); !ok {
			return Interface{}, fmt.Errorf("%s is not a named type", intfName)
		}

		intfType, ok := intfObj.Type().Underlying().(*types.Interface)
		if !ok {
			return Interface{}, fmt.Errorf("%s is not an interface", intfName)
		}

		// Get the file's imports:
		fileImprts := fileImports[fset.File(intfObj.Pos()).Pos(0)]

		intf := Interface{
			PackageName:   pkg.Name(),
			InterfaceName: intfObj.Name(),
		}
		for i := 0; i < intfType.NumMethods(); i++ {
			methodObj := intfType.Method(i)
			sig, ok := methodObj.Type().(*types.Signature)
			if !ok {
				log.Fatal("Method type is not a signature")
			}

			method := Method{
				Name: methodObj.Name(),
			}

			// Keep track of each of the method's parameters:
			paramsTuple := sig.Params()
			for j := 0; j < paramsTuple.Len(); j++ {
				paramObj := paramsTuple.At(j)
				param := Param{
					Name: paramObj.Name(),
					Type: types.TypeString(paramObj.Type(), Qualify(pkg, fileImprts, &intf.Imports)),
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
					Type: types.TypeString(resultObj.Type(), Qualify(pkg, fileImprts, &intf.Imports)),
				}
				method.Results = append(method.Results, result)
			}

			intf.Methods = append(intf.Methods, method)
		}

		return intf, nil
	}

	return Interface{}, errors.New("Interface not found")
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
				// If the package was not renamed in the import
				// statment, return the package's name:
				if imprt.Name == "" {
					other.Name()
				}
				// If the package was brought into this package
				// in an unqualified manner, don't qualify it:
				if imprt.Name == "." {
					return ""
				}
				// If the package was only imported for its
				// side-effects, skip over it, because it's
				// basically as if it was not imported at all:
				if imprt.Name == "_" {
					log.Println("Attempt to use type from package imported with the blank identifier")
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

				// Otherwise, the import was renamed, so return its name:
				return imprt.Name
			}
		}
		log.Printf("Package not found in file imports: '%s'", other.Path())
		return other.Name()
	}
}
