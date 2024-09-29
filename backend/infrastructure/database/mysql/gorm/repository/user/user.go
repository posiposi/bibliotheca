package user

import (
	"github.com/posiposi/project/backend/domain/model"
	ormModel "github.com/posiposi/project/backend/infrastructure/database/mysql/gorm/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) GetList() (users []model.User, err error) {
	dbUsers := []ormModel.User{}
	result := r.db.Find(&dbUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, u := range dbUsers {
		user := model.NewUser(u.Id, u.Name, u.Email, u.Password)
		users = append(users, *user)
	}
	return users, nil
}
