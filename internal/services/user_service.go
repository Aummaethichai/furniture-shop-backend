package services

import (
	"errors"
	"furniture-shop-backend/internal/dto"
	model "furniture-shop-backend/internal/models"
	repos "furniture-shop-backend/internal/repositories"
)

type UserService struct {
	repo repos.UserRepository // ✅ interface ไม่ต้องมี *
}

// ✅ รับ interface ตรง ๆ
func NewUserService(r repos.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	user := model.User{
		// ID:       uuid.New().String(),
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Role:     req.Role,
	}

	if err := s.repo.Insert(user); err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}, nil
}

func (s *UserService) FindUserByID(id string) (dto.UserResponse, error) {
	data, err := s.repo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if data == nil {
		return dto.UserResponse{}, errors.New("user not found")
	}
	return dto.UserResponse{
		ID:    data.ID,
		Email: data.Email,
		Name:  data.Name,
		Role:  data.Role,
	}, nil
}
