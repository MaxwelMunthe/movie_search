package gateway

import (
	"testing"

	"errors"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	expected := &ServiceImpl{}
	actual := NewService()
	assert.Equal(t, expected, actual)
}

func TestServiceImpl_ValidateParameter_Success(t *testing.T) {
	pagination := "2"

	service := &ServiceImpl{}
	err := service.validateParameter(pagination)
	assert.NoError(t, err)
}

func TestServiceImpl_ValidateParameter_Error_Invalid_Type(t *testing.T) {
	pagination := "a"
	expected := errors.New("pagination must be integer")

	service := &ServiceImpl{}
	err := service.validateParameter(pagination)
	assert.Equal(t, expected, err)
}

func TestServiceImpl_ValidateParameter_Error_Out_Range(t *testing.T) {
	pagination := "101"
	expected := errors.New("pagination must be 1-100")

	service := &ServiceImpl{}
	err := service.validateParameter(pagination)
	assert.Equal(t, expected, err)
}