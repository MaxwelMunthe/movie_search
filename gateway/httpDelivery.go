package gateway

import (
	"net/http"

	"movie-search/shared"
)

type HttpDelivery struct {
	service Service
}

func NewHttpDelivery(service Service) *HttpDelivery {
	return &HttpDelivery{service: service}
}

func (delivery HttpDelivery) GetMovieList(w http.ResponseWriter, r *http.Request) {
	pagination := r.URL.Query().Get("pagination")
	searchWord := r.URL.Query().Get("searchword")

	res, status, err := delivery.service.ProcessGetMovie(pagination, searchWord)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}
