// internal/services/user_service.go
package services

import "furniture-shop-backend/internal/repositories"

type UserService interface {
	GetUser(id string) (*repos.User, error)
}

type userService struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id string) (*repos.User, error) {
	// business logic เช่น ตรวจสิทธิ์, แปลงข้อมูล
	return s.repo.FindByID(id)
}
