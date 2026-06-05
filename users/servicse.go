package users

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/utils"
)

type UserService struct {
	User UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		User: userRepo,
	}
}

func (us *UserService) CreateUser(user UserFromRequest) error {
	err := user.Validate("")
	if err != nil {
		slog.Error("failed to validate user", "error", err)
		return err
	}
	bPass, err := utils.PassHash(user.Password)
	if err != nil {
		slog.Error("failed to hash password", "error", err)

		return err
	}
	err = us.User.Create(User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(bPass),
	})
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}
	return nil
}

func (us *UserService) Login(user UserFromRequest) (*User, error) {
	pass, err := us.User.GetPasswordByEmail(user.Email)
	if err != nil {
		slog.Error("failed to get user by email", "error", err)
		return nil, err
	}
	err = utils.ComparePass(user.Password, pass)
	if err != nil {
		slog.Error("failed to compare password", "error", err)
		return nil, err
	}
	userInfo, err := us.User.GetByEmail(user.Email)
	if err != nil {
		slog.Error("failed to get user by email", "error", err)
		return nil, err
	}
	return userInfo, nil
}

// func (us *UserService) User() error {
// 	return nil
// }

//	func (us *UserService) User() error {
//		return nil
//	}
