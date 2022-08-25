package user_repo_pg

import (
	"context"
	"database/sql"
	"time"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/ihsanbudiman/notes_app/sqlcpg"
	"gopkg.in/guregu/null.v4"
)

type postgresUserRepo struct {
	Source sqlcpg.Querier
}

// FindUserByUsernameOrEmailOrPhoneNumber implements domain.UserRepo
func (p postgresUserRepo) FindUserByUsernameOrEmailOrPhoneNumber(ctx context.Context, username string, email string, phone_number string) (domain.User, error) {
	data, err := p.Source.FindUserByUsernameOrEmailOrPhoneNumber(ctx, sqlcpg.FindUserByUsernameOrEmailOrPhoneNumberParams{
		Username: username,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		PhoneNumber: sql.NullString{
			String: phone_number,
			Valid:  true,
		},
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// FindUserByEmail implements domain.UserRepo
func (p postgresUserRepo) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	data, err := p.Source.FindUserByEmail(ctx, sql.NullString{
		String: email,
		Valid:  true,
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// FindUserByPhoneNumber implements domain.UserRepo
func (p postgresUserRepo) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	data, err := p.Source.FindUserByPhoneNumber(ctx, sql.NullString{
		String: phoneNumber,
		Valid:  true,
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// FindUserByUsername implements domain.UserRepo
func (p postgresUserRepo) FindUserByUsername(ctx context.Context, username string) (domain.User, error) {
	data, err := p.Source.FindUserByUsername(ctx, username)

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// GetUsers implements domain.UserRepo
func (p postgresUserRepo) GetUsers(ctx context.Context) ([]domain.User, error) {
	data, err := p.Source.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, v := range data {
		users = append(users, domain.User{
			ID:          int(v.ID),
			Name:        v.Name,
			Username:    v.Username,
			Password:    v.Password,
			Email:       null.String{NullString: v.Email},
			PhoneNumber: null.String{NullString: v.PhoneNumber},
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return users, nil
}

// FindUser implements domain.UserRepo
func (p postgresUserRepo) FindUser(ctx context.Context, id int) (domain.User, error) {
	data, err := p.Source.FindUser(ctx, int32(id))

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// Login implements domain.UserRepo
func (p postgresUserRepo) Login(ctx context.Context, username string) (domain.User, error) {
	data, err := p.Source.FindUserByUsername(ctx, username)

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// Register implements domain.UserRepo
func (p postgresUserRepo) Register(ctx context.Context, user domain.User) (domain.User, error) {
	// get sql.NullString from null.String

	data, err := p.Source.Register(ctx, sqlcpg.RegisterParams{
		Name:     user.Name,
		Username: user.Username,
		Email: sql.NullString{
			String: user.Email.String,
			Valid:  user.Email.Valid,
		},
		PhoneNumber: sql.NullString{
			String: user.PhoneNumber.String,
			Valid:  user.PhoneNumber.Valid,
		},
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          int(data.ID),
		Name:        data.Name,
		Username:    data.Username,
		Password:    data.Password,
		Email:       null.String{NullString: data.Email},
		PhoneNumber: null.String{NullString: data.PhoneNumber},
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

func NewPostgresUserRepo(source sqlcpg.Querier) domain.UserRepo {
	return &postgresUserRepo{source}
}
