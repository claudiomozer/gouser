package types

import (
	"regexp"
	"strings"

	"github.com/claudiomozer/gouser/internal/domain/err"
)

const (
	emailRegex = `^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
)

type Email string

func (e Email) Validate() error {
	regex := regexp.MustCompile(emailRegex)
	lowerEmail := strings.ToLower(e.String())

	if !regex.MatchString(lowerEmail) {
		return err.InvalidField("email")
	}
	return nil
}

func (e Email) String() string {
	return string(e)
}
