package services

import (
	"errors"
	"myginapp/models"
	"myginapp/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

// GetAllUsers fetches all users from the repository
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAllUsers()
}

func (s *userService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
