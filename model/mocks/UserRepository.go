package mock

import (
	"context"
	"errors"
	"project/model"
	"strconv"

	mock "github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) GetSingleUser(id int) (*model.User, error) {
	ret := m.Called(id)
	if ret.Get(0) == nil {
		return nil, errors.New("Unexpected")
	} else {
		result := ret.Get(0).(model.User)
		return &result, nil
	}
}

func (m *UserRepository) FindOne(ctx context.Context, username string, pass string) *model.User {
	ret := m.Called(ctx, username, pass)
	if ret.Get(0) == nil {
		return nil
	} else {
		result := ret.Get(0).(model.User)
		return &result
	}
}

func (m *UserRepository) UpdateDataSchool(user *model.School, id string) (*model.School, error) {
	id_int, _ := strconv.Atoi(id)
	ret := m.Called(user, id_int)

	var r0 *model.School
	if rf, ok := ret.Get(0).(func(*model.School, int) *model.School); ok {
		r0 = rf(user, id_int)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.School)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.School, int) error); ok {
		r1 = rf(user, id_int)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
