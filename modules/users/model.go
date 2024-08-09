package users

import (
	"GameBuy/helpers/common"
	"GameBuy/modules/role"
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	RoleId     string    `json:"role_id"`
	Role       role.Role `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if common.IsEmptyField(l.Username) {
		return errors.New("username is required")
	}

	if common.IsEmptyField(l.Password) {
		return errors.New("password is required")
	}

	return
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
}

func (s *SignUpRequest) ValidateSignUp() (err error) {
	if common.IsEmptyField(s.Username) {
		return errors.New("username is required")
	}

	if common.IsEmptyField(s.Password) {
		return errors.New("password is required")
	}
	if common.IsEmptyField(s.ReTypePassword) {
		return errors.New("retype password is required")
	}
	if s.Password != s.ReTypePassword {
		return errors.New("password and retype password not match")
	}

	re := regexp.MustCompile(`^(.{5,})$`)
	if !re.MatchString(s.Password) {
		return errors.New("please make sure that the password contains at least 5 character")
	}

	return nil
}

// set nilai untuk struct User
func (s *SignUpRequest) ConvertToModelForSignUp() (user User, err error) {
	hashedPassword, err := common.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password failed")
		return
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	return User{
		Username:   s.Username,
		Password:   hashedPassword,
		CreatedAt:  defaultField.CreatedAt,
		CreatedBy:  defaultField.CreatedBy,
		ModifiedAt: defaultField.ModifiedAt,
		ModifiedBy: defaultField.ModifiedBy,
	}, nil
}

type SignUpResponse struct {
}
