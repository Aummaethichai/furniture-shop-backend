// internal/repos/user_repo.go
package repos

import (
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"type:uuid;primaryKey"`
	Email string
	Name  string
}

type UserRepository interface {
	FindByID(id string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id string) (*User, error) {
	var u User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}