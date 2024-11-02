package services

import "tp/internal/models"

type UserService interface {
    CreateUser(user models.User) (models.User, error)
    GetUserByID(id int) (models.User, error)
}

type userService struct {
    userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
    return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id int) (models.User, error) {
    return s.userRepo.FindByID(id)
}

