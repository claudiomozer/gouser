package types

import (
	"strings"

	"github.com/claudiomozer/gouser/internal/domain/err"
)

const (
	minNameParts int = 2
	minNameSize  int = 3
)

type Name string

func (n Name) Validate() error {
	name_parts := strings.Split(n.String(), " ")

	if len(name_parts) < minNameParts {
		return err.InvalidField("name")
	}

	for _, n := range name_parts {
		if len(n) < minNameSize {
			return err.InvalidField("name")
		}
	}

	return nil
}

func (n Name) String() string {
	return string(n)
}
