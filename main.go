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
	intf, err := GetIntf(*dir, intfName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(intf)
}

type Intf struct {
	Name    string
	Methods []Method
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

func GetIntf(dir, intfName string) (Intf, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return Intf{}, fmt.Errorf("Erroring parsing file: %s", err)
	}

	conf := types.Config{Importer: importer.Default()}
	for pkgPath, pkgAST := range pkgs {
		files := []*ast.File{}
		for _, file := range pkgAST.Files {
			files = append(files, file)
		}

		pkg, err := conf.Check(pkgPath, fset, files, nil)
		if err != nil {
			return Intf{}, fmt.Errorf("Type error: %s", err)
		}

		intfObj := pkg.Scope().Lookup(intfName)
		if intfObj == nil {
			continue // Interface not found in this pkg, try next
		}

		if _, ok := intfObj.(*types.TypeName); !ok {
			return Intf{}, fmt.Errorf("%s is not a named type", intfName)
		}

		intfType, ok := intfObj.Type().Underlying().(*types.Interface)
		if !ok {
			return Intf{}, fmt.Errorf("%s is not an interface", intfName)
		}

		intf := Intf{
			Name: intfObj.Name(),
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

			paramsType := sig.Params()
			for j := 0; j < paramsType.Len(); j++ {
				paramObj := paramsType.At(j)
				param := Param{
					Name: paramObj.Name(),
					Type: paramObj.Type().String(),
				}
				method.Params = append(method.Params, param)
			}
			if len(method.Params) > 0 && sig.Variadic() {
				method.Params[len(method.Params)-1].Variadic = true
			}

			resultsType := sig.Results()
			for j := 0; j < resultsType.Len(); j++ {
				resultObj := resultsType.At(j)
				result := Result{
					Name: resultObj.Name(),
					Type: resultObj.Type().String(),
				}
				method.Results = append(method.Results, result)
			}

			intf.Methods = append(intf.Methods, method)
		}

		return intf, nil
	}

	return Intf{}, errors.New("Interface not found")
}
