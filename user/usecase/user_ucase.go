package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/ihsanbudiman/notes_app/helpers"
)

type UserUseCaseImpl struct {
	UserRepo domain.UserRepo
}

// CheckUniqueUserByEmail implements domain.UserUsecase
func (u UserUseCaseImpl) CheckUniqueUserByEmail(ctx context.Context, email string) (bool, error) {

	// check if email is not empty
	if email == "" {
		return false, errors.New("email cannot be empty")
	}

	// call repository
	user, err := u.UserRepo.FindUserByEmail(ctx, email)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// check if user is empty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

// CheckUniqueUserByPhoneNumber implements domain.UserUsecase
func (u UserUseCaseImpl) CheckUniqueUserByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {

	// check if phone number is not empty
	if phoneNumber == "" {
		return false, errors.New("phone number cannot be empty")
	}

	// call repository
	user, err := u.UserRepo.FindUserByPhoneNumber(ctx, phoneNumber)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// check if user is empty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// CheckUniqueUserByUsername implements domain.UserUsecase
func (u UserUseCaseImpl) CheckUniqueUserByUsername(ctx context.Context, username string) (bool, error) {

	// check if username is not empty
	if username == "" {
		return false, errors.New("username cannot be empty")
	}

	// call repository
	user, err := u.UserRepo.FindUserByUsername(ctx, username)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// check if user is empty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// FindUser implements domain.UserUsecase
func (u UserUseCaseImpl) FindUser(ctx context.Context, id int) (domain.User, error) {
	// check if id is not empty
	if id == 0 {
		return domain.User{}, errors.New("id cannot be empty")
	}

	// call repository
	user, err := u.UserRepo.FindUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Login implements domain.UserUsecase
func (u UserUseCaseImpl) Login(ctx context.Context, username string, password string) (domain.LoginResponse, error) {
	// check if username and password is not empty
	if username == "" || password == "" {
		return domain.LoginResponse{}, errors.New("username and password cannot be empty")
	}

	// call repository
	user, err := u.UserRepo.Login(ctx, username)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	fmt.Println(user.Password, password)

	// verify the password
	ok, err := helpers.ArgonVerify(password, user.Password)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	if !ok {
		return domain.LoginResponse{}, errors.New("password not match")
	}

	token, err := helpers.GenerateJwt(user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	finalData := domain.LoginResponse{
		Token: token,
		User:  user,
	}
	return finalData, nil
}

// Register implements domain.UserUsecase
func (u UserUseCaseImpl) Register(ctx context.Context, user domain.User) (domain.User, error) {
	// check if user is not empty
	if user.Name == "" || user.Username == "" || user.Password == "" {
		return domain.User{}, errors.New("name, username and password cannot be empty")
	}

	// check username, email, phone_number
	checkUser, err := u.UserRepo.FindUserByUsernameOrEmailOrPhoneNumber(ctx, user.Username, user.Email.ValueOrZero(), user.PhoneNumber.ValueOrZero())
	if err != nil && err != sql.ErrNoRows {
		return domain.User{}, err
	}

	if checkUser.ID != 0 {
		return domain.User{}, errors.New("username, email or phone number already exist")
	}

	// hash the password
	password, err := helpers.ArgonHash(user.Password)
	fmt.Println(password)
	if err != nil {
		return domain.User{}, err
	}

	// set the password
	user.Password = password

	// call repository
	user, err = u.UserRepo.Register(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func NewUserUseCase(ur domain.UserRepo) domain.UserUsecase {
	return &UserUseCaseImpl{
		UserRepo: ur,
	}
}
