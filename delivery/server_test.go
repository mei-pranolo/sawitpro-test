package delivery

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	instance := NewServer(NewServerOptions{Repository: &repository.Repository{}, Usecase: &usecase.Usecase{}})
	assert.NotNil(t, instance)
}
