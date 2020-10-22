package gateway

import (
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s ServiceMock) ProcessGetMovie(pagination, searchword string) (result interface{}, errStatus int, err error) {
	args := s.Called(pagination, searchword)
	return args.Get(0).(interface{}), args.Int(1), args.Error(2)
}
