package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnARole_WhenValidStringRoleIsProvided(t *testing.T) {
	testCases := []struct {
		name         string
		strRole      StringRole
		expectedRole Role
	}{
		{
			name:         "admin_role",
			strRole:      StringRole(admin),
			expectedRole: Admin,
		},
		{
			name:         "modifier_role",
			strRole:      StringRole(modifier),
			expectedRole: Modifier,
		},
		{
			name:         "watcher_role",
			strRole:      StringRole(watcher),
			expectedRole: Watcher,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			role := FromStringRole(testCase.strRole)
			assert.Equal(t, role, testCase.expectedRole)
		})
	}
}

func Test_ShouldReturnWatcherRole_WhenInvalidStringIsProvided(t *testing.T) {
	strRole := StringRole("invalid")
	role := FromStringRole(strRole)
	assert.Equal(t, role, Watcher)
}
