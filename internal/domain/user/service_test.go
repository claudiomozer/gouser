package user_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/types"
	"github.com/claudiomozer/gouser/internal/domain/user"
	"github.com/claudiomozer/gouser/mocks"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testArrage struct {
	service    *user.Service
	repository *mocks.MockUserRepository
}

func arrange(ctrl *gomock.Controller) testArrage {
	userRepository := mocks.NewMockUserRepository(ctrl)
	return testArrage{
		repository: userRepository,
		service:    user.NewService(userRepository),
	}
}

func Test_ShouldReturnErrorOnCreation_WhenRequiredsFieldsAreMissing(t *testing.T) {
	testCases := []struct {
		testName             string
		request              *user.CreateRequest
		expectedMissingField string
	}{
		{
			testName:             "missing_name_field",
			request:              &user.CreateRequest{},
			expectedMissingField: "name",
		}, {
			testName: "missing_email_field",
			request: &user.CreateRequest{
				Name: "some name",
			},
			expectedMissingField: "email",
		}, {
			testName: "missing_role_field",
			request: &user.CreateRequest{
				Name:  "some name",
				Email: "some@email.com",
			},
			expectedMissingField: "role",
		},
	}

	service := user.NewService(nil)
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			createErr := service.Create(context.TODO(), testCase.request)

			var domainErr *err.Error
			assert.NotNil(t, createErr)
			assert.ErrorAs(t, createErr, &domainErr)
			assert.Equal(t, domainErr.Code(), err.ErrMissingRequiredField)
			assert.Contains(t, domainErr.Message(), testCase.expectedMissingField)
		})
	}
}

func Test_ShouldReturnErrorOnCreation_WhenInvalidFieldsAreProvided(t *testing.T) {
	testCases := []struct {
		testName              string
		request               *user.CreateRequest
		expectingInvalidField string
	}{
		{
			testName: "invalid_name_field",
			request: &user.CreateRequest{
				Name:  "invalidname",
				Email: "some@email.com",
				Role:  "watcher",
			},
			expectingInvalidField: "name",
		}, {
			testName: "invalid_email_field",
			request: &user.CreateRequest{
				Name:  "some name",
				Email: "some_invalid_email.com",
				Role:  "watcher",
			},
			expectingInvalidField: "email",
		}, {
			testName: "invalid_role_field",
			request: &user.CreateRequest{
				Name:  "some name",
				Email: "some@email.com",
				Role:  "invalid",
			},
			expectingInvalidField: "role",
		},
	}

	service := user.NewService(nil)
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			createErr := service.Create(context.TODO(), testCase.request)

			var domainErr *err.Error
			assert.NotNil(t, createErr)
			assert.ErrorAs(t, createErr, &domainErr)
			assert.Equal(t, domainErr.Code(), err.ErrInvalidField)
			assert.Contains(t, domainErr.Message(), testCase.expectingInvalidField)
		})
	}
}

func Test_ShouldReturnError_WhenRepositoryReturnsFails(t *testing.T) {
	var (
		ctrl         = gomock.NewController(t)
		data         = arrange(ctrl)
		expectedUser = &user.Entity{
			Name:  "valid name",
			Email: "valid@email.com",
			Role:  types.Admin,
		}
	)
	defer ctrl.Finish()

	data.repository.EXPECT().Create(gomock.Any(), &userMatcher{
		entity: expectedUser,
	}).Return(errors.New("unexpected error"))
	err := data.service.Create(context.TODO(), &user.CreateRequest{
		Name:  expectedUser.Name,
		Email: expectedUser.Email,
		Role:  "admin",
	})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected error")
}

type userMatcher struct {
	entity *user.Entity
}

func (um *userMatcher) Matches(x any) bool {
	if another, ok := x.(*user.Entity); ok {
		return cmp.Equal(um.entity, another, cmpopts.IgnoreFields(user.Entity{}, "ID"))
	}
	return false
}

func (um *userMatcher) String() string {
	return fmt.Sprintf("user does not match: %v", um.entity)
}
