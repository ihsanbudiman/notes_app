package mock

import (
	"context"

	"github.com/ihsanbudiman/notes_app/sqlcpg"
	"github.com/stretchr/testify/mock"
)

type QuerierMock struct {
	mock.Mock
}

func (m *QuerierMock) FindUser(ctx context.Context, id int32) (sqlcpg.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(sqlcpg.User), args.Error(1)
}

func (m *QuerierMock) Login(ctx context.Context, params sqlcpg.LoginParams) (sqlcpg.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sqlcpg.User), args.Error(1)
}

func (m *QuerierMock) Register(ctx context.Context, user sqlcpg.User) (sqlcpg.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(sqlcpg.User), args.Error(1)
}

func (m *QuerierMock) GetUsers(ctx context.Context) ([]sqlcpg.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]sqlcpg.User), args.Error(1)
}
