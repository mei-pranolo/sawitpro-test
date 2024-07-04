// This file contains the repository implementation layer.

package repository

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	instance := NewRepository(NewRepositoryOptions{Dsn: "aaaa"})
	assert.NotNil(t, instance)
}
