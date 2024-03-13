package types

import (
	"strings"

	"github.com/claudiomozer/gouser/internal/domain/err"
)

type StringRole string

const (
	modifier string = "MODIFIER"
	watcher  string = "WATCHER"
	admin    string = "ADMIN"
)

func (s StringRole) Validate() error {
	switch strings.ToUpper(s.String()) {
	case admin, modifier, watcher:
		return nil
	default:
		return err.InvalidField("role")
	}
}

func (s StringRole) String() string {
	return strings.ToUpper(string(s))
}
