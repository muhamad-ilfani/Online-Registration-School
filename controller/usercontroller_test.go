package controller_test

import (
	"errors"
	"project/model"
	mocks "project/model/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_Get(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &model.User{
		Id:       1,
		Name:     "ilfani",
		Email:    "ilfani@gmail.com",
		Password: "ilfani",
		Role:     "admin",
		Ip:       1,
		SchoolId: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetSingleUser", 1).Return(*mockUser, nil).Once()
		a, err := mockUserRepo.GetSingleUser(1)
		assert.NoError(t, err)
		assert.NotNil(t, a)
		assert.Equal(t, a.Name, "ilfani")
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("failed", func(t *testing.T) {
		mockUserRepo.On("GetSingleUser", 99).Return(nil, errors.New("Unexpected")).Once()
		a, err := mockUserRepo.GetSingleUser(99)
		assert.NotNil(t, err)
		assert.Nil(t, a)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserController_Update(t *testing.T) {
	mockSchoolRepo := new(mocks.UserRepository)
	mockSchool := &model.School{
		Id:       1,
		Name:     "SMA N 1 Tegal",
		Capacity: 20,
		Students: []model.User{},
	}
	schoolID := mock.Anything
	mockEmpty := &model.School{}
	t.Run("success", func(t *testing.T) {
		mockSchoolRepo.On("UpdateDataSchool", mock.Anything, mock.Anything).Return(mockSchool, nil).Once()
		a, err := mockSchoolRepo.UpdateDataSchool(mockSchool, schoolID)
		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockSchoolRepo.AssertExpectations(t)
	})
	t.Run("failed", func(t *testing.T) {
		mockSchoolRepo.On("UpdateDataSchool", mock.Anything, mock.Anything).Return(mockEmpty, errors.New("Unexpected")).Once()
		a, err := mockSchoolRepo.UpdateDataSchool(mockSchool, schoolID)
		assert.Error(t, err)
		assert.Equal(t, mockEmpty, a)

		mockSchoolRepo.AssertExpectations(t)
	})
}
