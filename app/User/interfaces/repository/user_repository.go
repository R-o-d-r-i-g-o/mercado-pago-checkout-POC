package repository

import (
	"code-space-backend-api/app/User/entities"
	"code-space-backend-api/app/User/interfaces/dto/output"
	"code-space-backend-api/app/User/interfaces/mapper"
	"code-space-backend-api/infra/database/models"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(ctx context.Context, user entities.UserDomain) error
	GetUserByEmail(ctx context.Context, email string) (output.UserOutputDTO, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user entities.UserDomain) error {
	tx := r.db.WithContext(ctx)

	userModel := user.ToModel()
	return tx.Table(models.USER_TABLE_NAME).
		Create(&userModel).
		Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (output.UserOutputDTO, error) {
	db := r.db.WithContext(ctx)
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return output.UserOutputDTO{}, err
	}

	return mapper.UserModelToOutputDTO(user), nil
}
