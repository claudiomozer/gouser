package types

import (
	"testing"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/stretchr/testify/assert"
)

func Fuzz_ShouldReturnError_WhenInvalidEmailIsProvided(f *testing.F) {
	f.Add("")
	f.Add("invalidemail")
	f.Add("@email.com")
	f.Add("email.com")
	f.Add("mail@mail")
	f.Add("email@@email.com")
	f.Add("invalid@email,com")
	f.Add("email @gmail.com")
	f.Add(".@gmail.com")

	f.Fuzz(func(t *testing.T, email string) {
		validateErr := Email(email).Validate()

		var domainErr *err.Error
		assert.NotNil(t, validateErr)
		assert.ErrorAs(t, validateErr, &domainErr)
		assert.Equal(t, domainErr.Code(), err.ErrInvalidField)
		assert.Contains(t, domainErr.Message(), "email")
	})
}

func Fuzz_ShouldNotReturnError_WhenValidEmailIsProvided(f *testing.F) {
	f.Add("valid@email.com")
	f.Add("valid_underscore@email.com")
	f.Add("valid.point@email.com")
	f.Add("valid@sub.domain.com")
	f.Add("valid12number@mail.com")

	f.Fuzz(func(t *testing.T, email string) {
		validateErr := Email(email).Validate()
		assert.Nil(t, validateErr)
	})
}
