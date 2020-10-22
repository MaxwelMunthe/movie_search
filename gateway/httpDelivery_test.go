package gateway

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpDelivery(t *testing.T) {
	expected := &HttpDelivery{}
	actual := NewHttpDelivery(nil)
	assert.Equal(t, expected, actual)
}

func TestHttpDelivery_GetMovieList_Success(t *testing.T) {
	serviceMock := &ServiceMock{}
	delivery := &HttpDelivery{
		service: serviceMock,
	}

	pagination := "2"
	searchWord := "batman"

	serviceMock.On("ProcessGetMovie", pagination, searchWord).Return(map[string]interface{}{}, http.StatusOK, nil)
	r := httptest.NewRequest(http.MethodGet, "/search?pagination="+pagination+"&searchword="+searchWord, nil)
	w := httptest.NewRecorder()
	delivery.GetMovieList(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHttpDelivery_GetMovieList_Error_500(t *testing.T) {
	serviceMock := &ServiceMock{}
	delivery := &HttpDelivery{
		service: serviceMock,
	}

	pagination := "2"
	searchWord := "batman"

	serviceMock.On("ProcessGetMovie", pagination, searchWord).Return(map[string]interface{}{}, http.StatusInternalServerError, errors.New("test error"))
	r := httptest.NewRequest(http.MethodGet, "/search?pagination="+pagination+"&searchword="+searchWord, nil)
	w := httptest.NewRecorder()
	delivery.GetMovieList(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
