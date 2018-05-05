package main

import (
	"fmt"
	"strings"
)

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

func (m *Method) ParamsString() string {
	var strs []string
	for _, param := range m.Params {
		typ := param.Type
		if param.Variadic {
			// Remove "[]" and replace with "..."
			typ = fmt.Sprintf("...%s", strings.TrimPrefix(param.Type, "[]"))
		}

		str := typ
		if param.Name != "" {
			str = fmt.Sprintf("%s %s", param.Name, typ)
		}
		strs = append(strs, str)
	}
	return strings.Join(strs, ", ")
}

func (m *Method) NamedParamsString() string {
	var strs []string
	for i, param := range m.Params {
		name := param.Name
		if name == "" || name == "_" {
			name = fmt.Sprintf("param%d", i+1)
		}

		typ := param.Type
		if param.Variadic {
			// Remove "[]" and replace with "..."
			typ = fmt.Sprintf("...%s", strings.TrimPrefix(param.Type, "[]"))
		}
		strs = append(strs, fmt.Sprintf("%s %s", name, typ))
	}
	return strings.Join(strs, ", ")
}

func (m *Method) ParamNamesString() string {
	var names []string
	for i, param := range m.Params {
		name := param.Name
		if name == "" || name == "_" {
			name = fmt.Sprintf("param%d", i+1)
		}
		if param.Variadic {
			name = fmt.Sprintf("%s...", name)
		}
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}

func (m *Method) ResultsString() string {
	var (
		strs  []string
		named bool
	)
	for _, result := range m.Results {
		str := result.Type
		if result.Name != "" {
			named = true
			str = fmt.Sprintf("%s %s", result.Name, result.Type)
		}
		strs = append(strs, str)
	}
	if len(strs) > 1 || named {
		return fmt.Sprintf("(%s)", strings.Join(strs, ", "))
	}
	return strings.Join(strs, ", ")
}
