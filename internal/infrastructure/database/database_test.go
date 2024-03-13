package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldPanic_WhenParseConfigFails(t *testing.T) {
	config := &ConnectionConfig{} // must fail without any variable for connection
	assert.Panics(t, func() {
		ConnectToDatabase(context.TODO(), config)
	})
}
