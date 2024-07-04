package usecase

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {
	instance := NewUsecase(&repository.Repository{})
	assert.NotNil(t, instance)
}
