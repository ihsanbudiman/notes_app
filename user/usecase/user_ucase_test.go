package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/ihsanbudiman/notes_app/user/repository/postgres/mock"
	mmock "github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v4"
)

func TestUserUseCaseImpl_CheckUniqueUserByEmail(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByEmail", ctx, "test").Return(domain.User{}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByEmail(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByEmail", ctx, "test").Return(domain.User{
			ID: 1,
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByEmail(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		// expected: false
		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// sql.ErrNoRows
	t.Run("sql.ErrNoRows", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByEmail", ctx, "test").Return(domain.User{}, sql.ErrNoRows)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByEmail(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	// error
	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByEmail", ctx, "test").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByEmail(ctx, "test")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// err empty email
	t.Run("err empty email", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByEmail", ctx, "").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByEmail(ctx, "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})
}

func TestUserUseCaseImpl_CheckUniqueUserByPhoneNumber(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "test").Return(domain.User{}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "test").Return(domain.User{
			ID: 1,
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		// expected: false
		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})
	// sql.ErrNoRows
	t.Run("sql.ErrNoRows", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "test").Return(domain.User{}, sql.ErrNoRows)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "test")
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	// error
	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "test").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "test")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// empty phone number
	t.Run("empty phone number", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// nil phone number
	t.Run("nil phone number", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByPhoneNumber", ctx, "").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByPhoneNumber(ctx, "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

}

func TestUserUseCaseImpl_CheckUniqueUserByUsername(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "test").Return(domain.User{}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "test").Return(domain.User{
			ID: 1,
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "test")

		if err != nil {
			t.Errorf("error: %v", err)
		}

		// expected: false
		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})
	// sql.ErrNoRows
	t.Run("sql.ErrNoRows", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "test").Return(domain.User{}, sql.ErrNoRows)
		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "test")
		if err != nil {
			t.Errorf("error: %v", err)
		}

		if !ok {
			t.Errorf("expected: true, got: %v", ok)
		}

	})

	// error
	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "test").Return(domain.User{}, errors.New("error"))
		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "test")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// empty username
	t.Run("empty username", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "").Return(domain.User{}, errors.New("error"))
		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})

	// nil username
	t.Run("nil username", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUserByUsername", ctx, "").Return(domain.User{}, errors.New("error"))
		useCase := UserUseCaseImpl{UserRepo: mock}

		ok, err := useCase.CheckUniqueUserByUsername(ctx, "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if ok {
			t.Errorf("expected: false, got: %v", ok)
		}

	})
}

func TestUserUseCaseImpl_FindUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUser", ctx, 1).Return(domain.User{
			ID: 1,
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		user, err := useCase.FindUser(ctx, 1)

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if user.ID != 1 {
			t.Errorf("expected: 1, got: %v", user.ID)
		}

	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUser", ctx, 1).Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		user, err := useCase.FindUser(ctx, 1)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if user.ID != 0 {
			t.Errorf("expected: 0, got: %v", user.ID)
		}

	})

	t.Run("empty id", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("FindUser", ctx, 0).Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		user, err := useCase.FindUser(ctx, 0)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

		if user.ID != 0 {
			t.Errorf("expected: 0, got: %v", user.ID)
		}

	})
}

func TestUserUseCaseImpl_Login(t *testing.T) {
	t.Run("empty username", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "", "aasdfds")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

	})

	t.Run("empty password", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}

	})

	t.Run("Err call repo login", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("Login", ctx, "test").Return(domain.User{}, errors.New("error"))

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "sdfds")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("Err argon verify", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("Login", ctx, "test").Return(domain.User{
			ID:       1,
			Password: "t=3,p=2$DqP3+EF5kXzKAUwuB+4Z4g$V9la36RAO3wdC0UzHPL2Jzl/dkvFhgDUsk+0kwS381M",
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "123")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("Err password do not match", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("Login", ctx, "test").Return(domain.User{
			ID:       1,
			Password: "$argon2id$v=19$m=65536,t=3,p=2$DqP3+EF5kXzKAUwuB+4Z4g$V9la36RAO3wdC0UzHPL2Jzl/dkvFhgDUsk+0kwS381M",
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "ihexsan123")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	// err generate jwt
	t.Run("Err generate jwt", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("Login", ctx, "test").Return(domain.User{
			ID:       0,
			Password: "$argon2id$v=19$m=65536,t=3,p=2$DqP3+EF5kXzKAUwuB+4Z4g$V9la36RAO3wdC0UzHPL2Jzl/dkvFhgDUsk+0kwS381M",
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "ihsan123")

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	// success
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mock := new(mock.MockPostgresUserRepo)
		mock.On("Login", ctx, "test").Return(domain.User{
			ID:       1,
			Password: "$argon2id$v=19$m=65536,t=3,p=2$DqP3+EF5kXzKAUwuB+4Z4g$V9la36RAO3wdC0UzHPL2Jzl/dkvFhgDUsk+0kwS381M",
		}, nil)

		useCase := UserUseCaseImpl{UserRepo: mock}

		_, err := useCase.Login(ctx, "test", "ihsan123")

		if err != nil {
			t.Errorf("expected: nil, got: %v", err)
		}
	})

}

func TestUserUseCaseImpl_Register(t *testing.T) {

	t.Run("empty username", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "ihsanbudiman",
			Username:    "",
			Email:       null.StringFrom(""),
			PhoneNumber: null.StringFrom(""),
			Password:    "ihsan123",
		}

		_, err := useCase.Register(ctx, user)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("empty password", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "ihsanbudiman",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom(""),
			PhoneNumber: null.StringFrom(""),
			Password:    "",
		}

		_, err := useCase.Register(ctx, user)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("empty name", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom(""),
			PhoneNumber: null.StringFrom(""),
			Password:    "ihsan123",
		}

		_, err := useCase.Register(ctx, user)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	// err in FindUserByUsernameOrEmailOrPhoneNumber
	t.Run("err in FindUserByUsernameOrEmailOrPhoneNumber", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "ihsanbudiman",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom(""),
			PhoneNumber: null.StringFrom(""),
			Password:    "ihsan123",
		}

		m.On("FindUserByUsernameOrEmailOrPhoneNumber", ctx, user.Username, user.Email.String, user.PhoneNumber.String).Return(domain.User{}, errors.New("error"))

		_, err := useCase.Register(ctx, user)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("user found", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "ihsanbudiman",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom(""),
			PhoneNumber: null.StringFrom(""),
			Password:    "ihsan123",
		}

		m.On("FindUserByUsernameOrEmailOrPhoneNumber", ctx, user.Username, user.Email.String, user.PhoneNumber.String).Return(domain.User{
			ID: 1,
		}, nil)

		_, err := useCase.Register(ctx, user)

		if err == nil {
			t.Errorf("expected: error, got: nil")
		}
	})

	t.Run("register err", func(t *testing.T) {
		m := new(mock.MockPostgresUserRepo)
		useCase := UserUseCaseImpl{UserRepo: m}
		ctx := context.Background()

		user := domain.User{
			Name:        "ihsanbudiman",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom("sdfg"),
			PhoneNumber: null.StringFrom("sdfg"),
			Password:    "ihsan",
		}

		m.On("FindUserByUsernameOrEmailOrPhoneNumber", ctx, user.Username, user.Email.String, user.PhoneNumber.String).Return(domain.User{}, nil)
		m.On("Register", ctx, domain.User{
			Name:        "ihsanbudiman",
			Username:    "ihsanbudiman",
			Email:       null.StringFrom("sdfg"),
			PhoneNumber: null.StringFrom("sdfg"),
			Password:    string(mmock.AnythingOfType("string")),
		}).Return(domain.User{}, nil)

		useCase.Register(ctx, user)
		// if err != nil {
		// 	t.Errorf("expected: nil, got: %v", err)
		// }

	})
}

func TestNewUserUseCase(t *testing.T) {
	m := new(mock.MockPostgresUserRepo)

	useCase := NewUserUseCase(m)

	if useCase == nil {
		t.Errorf("expected: useCase, got: nil")
	}

}
