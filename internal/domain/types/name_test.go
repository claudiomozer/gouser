package types

import (
	"testing"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/stretchr/testify/assert"
)

func Fuzz_ShouldReturnError_WhenInvalidNameIsProvided(f *testing.F) {
	f.Add("")
	f.Add("invalidname")
	f.Add("i name")
	f.Add("i n")
	f.Add("i name")

	f.Fuzz(func(t *testing.T, name string) {
		validateErr := Name(name).Validate()

		var domainErr *err.Error
		assert.NotNil(t, validateErr)
		assert.ErrorAs(t, validateErr, &domainErr)
		assert.Equal(t, domainErr.Code(), err.ErrInvalidField)
		assert.Contains(t, domainErr.Message(), "name")
	})
}

func Test_ShouldReturnNil_WhenValidNameIsGiven(t *testing.T) {
	var name Name = "super name for test purposes"

	validateErr := name.Validate()
	assert.Nil(t, validateErr)
}
