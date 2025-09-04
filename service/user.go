package service

import (
	"go-project/model"
	"go-project/repository"
)

type UserService interface {
	Create(user model.User) (uint64, error)
	FindByNickOrName(nickOrName string) ([]model.User, error)
	FindByMail(mail string) (model.User, error)
	FindById(ID uint64) (model.User, error)
	FindAll() ([]model.User, error)
	Update(ID uint64, user model.User) error
	Delete(ID uint64) error
}

type userServiceImpl struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return userServiceImpl{repo}
}

// Create implements UserService.
func (u userServiceImpl) Create(user model.User) (uint64, error) {
	return u.repo.Create(user)
}

// Delete implements UserService.
func (u userServiceImpl) Delete(ID uint64) error {
	return u.repo.Delete(ID)
}

// FindAll implements UserService.
func (u userServiceImpl) FindAll() ([]model.User, error) {
	return u.repo.FindAll()
}

// FindById implements UserService.
func (u userServiceImpl) FindById(ID uint64) (model.User, error) {
	return u.repo.FindByID(ID)
}

// FindByMail implements UserService.
func (u userServiceImpl) FindByMail(mail string) (model.User, error) {
	return u.repo.FindByMail(mail)
}

// FindByNickOrName implements UserService.
func (u userServiceImpl) FindByNickOrName(nickOrName string) ([]model.User, error) {
	return u.repo.FindByNickOrName(nickOrName)
}

// Update implements UserService.
func (u userServiceImpl) Update(ID uint64, user model.User) error {
	return u.repo.Update(ID, user)
}
