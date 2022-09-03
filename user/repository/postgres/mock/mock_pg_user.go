package mock

import (
	"context"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/stretchr/testify/mock"
)

type MockPostgresUserRepo struct {
	mock.Mock
}

// FindUser implements domain.UserRepo
func (m *MockPostgresUserRepo) FindUser(ctx context.Context, id int) (domain.User, error) {
	// mock
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

// FindUserByEmail implements domain.UserRepo
func (m *MockPostgresUserRepo) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	// mock
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}

// FindUserByPhoneNumber implements domain.UserRepo
func (m *MockPostgresUserRepo) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	// mock
	args := m.Called(ctx, phoneNumber)
	return args.Get(0).(domain.User), args.Error(1)
}

// FindUserByUsername implements domain.UserRepo
func (m *MockPostgresUserRepo) FindUserByUsername(ctx context.Context, username string) (domain.User, error) {
	// mock
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}

// FindUserByUsernameOrEmailOrPhoneNumber implements domain.UserRepo
func (m *MockPostgresUserRepo) FindUserByUsernameOrEmailOrPhoneNumber(ctx context.Context, username string, email string, phone_number string) (domain.User, error) {
	// mock
	args := m.Called(ctx, username, email, phone_number)
	return args.Get(0).(domain.User), args.Error(1)
}

// Login implements domain.UserRepo
func (m *MockPostgresUserRepo) Login(ctx context.Context, username string) (domain.User, error) {
	// mock
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}

// Register implements domain.UserRepo
func (m *MockPostgresUserRepo) Register(ctx context.Context, user domain.User) (domain.User, error) {
	// mock
	args := m.Called(ctx, user)
	return args.Get(0).(domain.User), args.Error(1)
}
