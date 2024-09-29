package user

import (
	"testing"

	"github.com/posiposi/project/backend/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (urm *UserRepositoryMock) GetList() ([]model.User, error) {
	r := urm.Called()
	return r.Get(0).([]model.User), r.Error(1)
}

func TestGetList(t *testing.T) {
	m := new(UserRepositoryMock)
	m.On("GetList").Return([]model.User{}, nil)
	testObj := NewUserUseCase(m)
	r, err := testObj.GetList()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	m.AssertExpectations(t)
}
