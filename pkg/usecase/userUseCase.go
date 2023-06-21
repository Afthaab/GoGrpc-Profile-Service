package usecase

import (
	"errors"

	"github.com/profile/service/pkg/domain"
	interfaces "github.com/profile/service/pkg/repository/interface"
	useCase "github.com/profile/service/pkg/usecase/interface"
	utility "github.com/profile/service/pkg/utils"
)

type UserRepo struct {
	Repo interfaces.UserRepo
}

func (u *UserRepo) ViewProfile(user domain.User) (domain.User, error) {
	user, err := u.Repo.FindProfile(user)
	if err != nil {
		return user, errors.New("User details not found")
	}
	return user, nil
}

func (u *UserRepo) EditProfile(user domain.User) error {
	err := u.Repo.EditProfile(user)
	if err == 0 {
		return errors.New("Could not update the user details")
	}
	return nil
}
func (u *UserRepo) ChangePassword(passwordData domain.Password) error {
	user := domain.User{
		Id: passwordData.Id,
	}
	// finding the userDetails through user id from middleware
	user, err := u.Repo.FindProfile(user)
	if err != nil {
		return errors.New("User details not found")
	}

	// checking the entered old passwords
	if !utility.VerifyPassword(passwordData.Oldpassword, user.Password) {
		return errors.New("Current Password did not match")
	}

	// Hash the new password
	passwordData.Newpassword = utility.HashPassword(passwordData.Newpassword)

	//updating the password
	result := u.Repo.UpdatePassword(passwordData)
	if result == 0 {
		return errors.New("Could not update the new Password")
	}

	return nil
}

func (u *UserRepo) AddAddress(addressData domain.Address) (domain.Address, error) {
	addressData, err := u.Repo.CreateAddress(addressData)
	if err != nil {
		return addressData, errors.New("Could not create the Address")
	}
	return addressData, nil
}

func NewUserUseCase(repo interfaces.UserRepo) useCase.UserUseCase {
	return &UserRepo{
		Repo: repo,
	}
}
