package iface

import (
	"fmt"
	"go/token"
	"strings"
)

type Interface struct {
	Name    string
	Package string
	Imports []Import
	Methods Methods
}

type Import struct {
	Name string
	Path string
}

func (i *Import) String() string {
	if i.Name != "" {
		return fmt.Sprintf("%s \"%s\"", i.Name, i.Path)
	}
	return fmt.Sprintf("\"%s\"", i.Path)
}

type Method struct {
	Name    string
	Params  Params
	Results Results

	pos token.Pos
}

type Methods []Method

func (m Methods) Len() int           { return len(m) }
func (m Methods) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Methods) Less(i, j int) bool { return m[i].pos < m[j].pos }

type Param struct {
	Name     string
	Type     string
	Variadic bool
}

func (p *Param) Named(i int) string {
	if p.Name == "" || p.Name == "_" {
		return fmt.Sprintf("param%d", i+1)
	}
	return p.Name
}

func (p *Param) FieldName(i int) string {
	return strings.Title(p.Named(i))
}

func (p *Param) String() string {
	if p.Name != "" {
		return fmt.Sprintf("%s %s", p.Name, p.TypeString())
	}
	return p.TypeString()
}

func (p *Param) TypeString() string {
	if p.Variadic {
		return fmt.Sprintf("...%s", strings.TrimPrefix(p.Type, "[]"))
	}
	return p.Type
}

type Params []Param

func (ps Params) String() string {
	var strs []string
	for _, p := range ps {
		strs = append(strs, p.String())
	}
	return strings.Join(strs, ", ")
}

func (ps Params) NamedString() string {
	var strs []string
	for i, p := range ps {
		strs = append(strs, fmt.Sprintf("%s %s", p.Named(i), p.TypeString()))
	}
	return strings.Join(strs, ", ")
}

func (ps Params) VariadicArgsString() string {
	var args []string
	for i, p := range ps {
		arg := p.Named(i)
		if p.Variadic {
			arg = fmt.Sprintf("%s...", arg)
		}
		args = append(args, arg)
	}
	return strings.Join(args, ", ")
}

func (ps Params) ArgsString() string {
	var args []string
	for i, p := range ps {
		args = append(args, p.Named(i))
	}
	return strings.Join(args, ", ")
}

func (ps Params) NamedFields() string {
	var strs []string
	for i, p := range ps {
		strs = append(strs, fmt.Sprintf("\t%s %s", p.FieldName(i), p.Type))
	}
	return strings.Join(strs, "\n")
}

type Result struct {
	Name string
	Type string
}

func (r *Result) String() string {
	if r.Name != "" {
		return fmt.Sprintf("%s %s", r.Name, r.Type)
	}
	return r.Type
}

func (r *Result) Named(i int) string {
	if r.Name == "" || r.Name == "_" {
		return fmt.Sprintf("result%d", i+1)
	}
	return r.Name
}

func (r *Result) FieldName(i int) string {
	return strings.Title(r.Named(i))
}

type Results []Result

func (rs Results) String() string {
	var (
		strs  []string
		named bool
	)
	for _, r := range rs {
		if r.Name != "" {
			named = true
		}
		strs = append(strs, r.String())
	}
	if len(strs) > 1 || named {
		return fmt.Sprintf("(%s)", strings.Join(strs, ", "))
	}
	return strings.Join(strs, ", ")
}

func (rs Results) NamedFields() string {
	var strs []string
	for i, r := range rs {
		strs = append(strs, fmt.Sprintf("\t%s %s", r.FieldName(i), r.Type))
	}
	return strings.Join(strs, "\n")
}

func (rs Results) ReturnList(structName string) string {
	var strs []string
	for i, r := range rs {
		strs = append(strs, fmt.Sprintf("%s.%s", structName, r.FieldName(i)))
	}
	return strings.Join(strs, ",")
}
