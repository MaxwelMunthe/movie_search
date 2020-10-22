package gateway

type Service interface {
	ProcessGetMovie(pagination, searchword string) (result interface{}, errStatus int, err error)
}
