package gateway

import (
	"github.com/go-chi/chi"
)

type Routes struct {
	handler *HttpDelivery
}

func NewRoutes(handler *HttpDelivery) *Routes {
	return &Routes{handler: handler}
}

func InjectRoutes() *Routes {
	services := NewService()
	delivery := NewHttpDelivery(services)

	return NewRoutes(delivery)
}

func (routes Routes) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/search", routes.handler.GetMovieList)
	})
}
