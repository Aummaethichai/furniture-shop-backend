// internal/repos/user_repo.go
package repos

import (
	"errors"  
	"furniture-shop-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (*model.User, error)
	Insert(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
    var u model.User
    if err := r.db.First(&u, "id = ?", id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) Insert(user model.User) error {
	return r.db.Create(&user).Error
}