package user

import (
	"testing"

	"github.com/posiposi/project/backend/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (uum *UserUsecaseMock) GetList() ([]model.User, error) {
	r := uum.Called()
	return r.Get(0).([]model.User), r.Error(1)
}

func TestGetUserList(t *testing.T) {
	uum := new(UserUsecaseMock)
	uum.On("GetList").Return([]model.User{}, nil)
	r, err := uum.GetList()
	assert.NoError(t, err)
	assert.Equal(t, []model.User{}, r)
}
