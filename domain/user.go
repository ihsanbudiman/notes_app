package domain

import (
	"context"
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Username    string      `json:"username"`
	Email       null.String `json:"email"`
	PhoneNumber null.String `json:"phone_number"`
	Password    string      `json:"password"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// empty password json marshal
func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""

	type Alias User
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&u),
	})
}

type UserRepo interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, username string) (User, error)
	FindUser(ctx context.Context, id int) (User, error)
	FindUserByUsername(ctx context.Context, username string) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	FindUserByUsernameOrEmailOrPhoneNumber(ctx context.Context, username, email, phone_number string) (User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, username string, password string) (User, error)
	FindUser(ctx context.Context, id int) (User, error)
	CheckUniqueUserByUsername(ctx context.Context, username string) (bool, error)
	CheckUniqueUserByEmail(ctx context.Context, email string) (bool, error)
	CheckUniqueUserByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
}
