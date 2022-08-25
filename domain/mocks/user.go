package mocks

import (
	"context"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

// Login implements domain.UserRepo
func (m *UserRepoMock) Login(ctx context.Context, username string, password string) (domain.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

// Register implements domain.UserRepo
func (m *UserRepoMock) Register(ctx context.Context, user domain.User) (domain.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(domain.User), args.Error(1)
}

// GetUsers implements domain.UserRepo
func (m *UserRepoMock) GetUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

// FindUser implements domain.UserRepo
func (m *UserRepoMock) FindUser(ctx context.Context, id int) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}
