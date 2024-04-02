package iface

import (
	"cmp"
	"fmt"
	"go/token"
	"strings"
)

type Interface struct {
	Name       string
	TypeParams TypeParams
	Package    string
	Imports    []Import
	Methods    Methods
}

type TypeParam struct {
	Name       string
	Constraint string
}

func (t TypeParam) String() string {
	return fmt.Sprintf("%s %s", t.Name, t.Constraint)
}

type TypeParams []TypeParam

func (t TypeParams) mapString(f func(TypeParam) string) string {
	if len(t) == 0 {
		return ""
	}
	var s strings.Builder
	s.WriteByte('[')
	for i, typeParam := range t {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(f(typeParam))
	}
	s.WriteByte(']')
	return s.String()
}

func (t TypeParams) String() string {
	return t.mapString(TypeParam.String)
}

func (t TypeParams) Names() string {
	return t.mapString(func(t TypeParam) string { return t.Name })
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

	// String representation of the interface explicitly requiring this method
	srcIface string
	pos      token.Pos
}

type Methods []Method

func (m Methods) Len() int      { return len(m) }
func (m Methods) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m Methods) Less(i, j int) bool {
	// Group methods by source interface. This grouping matters when the mocked
	// interface comprises interfaces defined in multiple files, in which case
	// the token positions alone don't have a stable ordering.
	switch cmp.Compare(m[i].srcIface, m[j].srcIface) {
	case -1: // less
		return true
	case 1: // greater
		return false
	default: // equal
		return m[i].pos < m[j].pos
	}
}

type Param struct {
	Name     string
	Type     string
	Variadic bool
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
		name := p.Name
		if name == "" || name == "_" {
			name = fmt.Sprintf("param%d", i+1)
		}

		strs = append(strs, fmt.Sprintf("%s %s", name, p.TypeString()))
	}
	return strings.Join(strs, ", ")
}

func (ps Params) ArgsString() string {
	var args []string
	for i, param := range ps {
		arg := param.Name
		if arg == "" || arg == "_" {
			arg = fmt.Sprintf("param%d", i+1)
		}
		if param.Variadic {
			arg = fmt.Sprintf("%s...", arg)
		}
		args = append(args, arg)
	}
	return strings.Join(args, ", ")
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
