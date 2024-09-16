package user

import "github.com/posiposi/project/backend/domain/model"

type UserRepository interface {
	GetList() ([]model.User, error)
}
