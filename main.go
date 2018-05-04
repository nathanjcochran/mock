package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
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
	dir := flag.String("d", ".", "Directory to search for interface in")
	outFile := flag.String("o", "", "Output file (default stdout)")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] interface\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// First argument is interface name:
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough args")
	}
	intfName := args[0]

	// Parse the package and get info about the interface:
	intf, err := GetInterface(*dir, intfName)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the template:
	tmpl, err := template.New("default").Parse(defaultTmpl)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	// Execute/output the template:
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, &intf); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	// Format it with go fmt:
	result, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Error formating output: %s", err)
	}

	// Open the file, if provided, or use stdout:
	out := os.Stdout
	if *outFile != "" {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatalf("Error creating output file: %s", err)
		}
		defer out.Close()
	}

	// Write the formatted output to the file:
	if _, err := out.Write(result); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
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
