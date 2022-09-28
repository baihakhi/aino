package service

import (
	"context"
	"errors"

	"auth/src/constanta/crypt"
	"auth/src/models"
	repository "auth/src/repositories"
	"auth/src/service/auth"
)

type authUC struct {
	repo repository.UserRepository
}

type Auth interface {
	Register(ctx context.Context, params auth.RegisterRequest) error
	Login(ctx context.Context, params auth.LoginRequest) (*models.User, error)
	Logout(ctx context.Context, email string) error
}

func NewAuthUC(r repository.UserRepository) Auth {
	return &authUC{
		repo: r,
	}
}

func (u *authUC) Register(ctx context.Context, params auth.RegisterRequest) error {
	var (
		encryptedPassword string
		err               error
		user              *models.User
	)

	if err = params.Validate(); err != nil {
		return err
	}

	user, _ = u.repo.GetUserByEmail(ctx, params.Email)

	if user != nil {
		return errors.New("email is used")
	}

	encryptedPassword, err = crypt.EncryptPassword(params.Password)
	if err != nil {
		return err
	}

	req := &models.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: encryptedPassword,
		IsLogin:  false,
	}

	err = u.repo.InsertUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *authUC) Login(ctx context.Context, params auth.LoginRequest) (*models.User, error) {

	var (
		err  error
		user *models.User
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	user, err = u.repo.GetUserByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, errors.New("email not registered")
	}

	err = crypt.ComparePassword(params.Password, user.Password)
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	err = u.repo.UpdateStatusLoginTrue(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	user, _ = u.repo.GetUserByEmail(ctx, params.Email)

	return user, nil
}

func (u *authUC) Logout(ctx context.Context, email string) error {

	var (
		err  error
		user *models.User
	)

	user, _ = u.repo.GetUserByEmail(ctx, email)
	if !user.IsLogin {
		return errors.New("you need to login first")
	}

	err = u.repo.UpdateStatusLoginFalse(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
