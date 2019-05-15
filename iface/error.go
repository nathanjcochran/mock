package iface

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/packages"
)

type TypeErrors struct {
	Errs []packages.Error
}

func (e *TypeErrors) Error() string {
	var strs []string
	for _, err := range e.Errs {
		strs = append(strs, err.Error())
	}
	return fmt.Sprintf("encountered type errors: \n%s", strings.Join(strs, "\n"))
}
