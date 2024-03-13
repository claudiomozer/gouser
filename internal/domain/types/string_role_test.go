package types

import (
	"testing"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/stretchr/testify/assert"
)

func Fuzz_ShouldReturnError_WhenInvalidRoleNameIsProvided(f *testing.F) {
	f.Add("")
	f.Add("invalid")

	f.Fuzz(func(t *testing.T, role string) {
		validateErr := StringRole(role).Validate()

		var domainErr *err.Error
		assert.NotNil(t, validateErr)
		assert.ErrorAs(t, validateErr, &domainErr)
		assert.Equal(t, domainErr.Code(), err.ErrInvalidField)
		assert.Contains(t, domainErr.Message(), "role")
	})
}

func Fuzz_ShouldNotReturnError_WhenValidRoleNameIsProvided(f *testing.F) {
	f.Add("admin")
	f.Add("ADMIN")
	f.Add("WATCHER")
	f.Add("MODIFIER")

	f.Fuzz(func(t *testing.T, role string) {
		validateErr := StringRole(role).Validate()
		assert.Nil(t, validateErr)
	})
}
