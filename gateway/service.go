package gateway

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"movie-search/proto"

	"google.golang.org/grpc"
)

type ServiceImpl struct {
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (s ServiceImpl) ProcessGetMovie(pagination, searchword string) (result interface{}, errStatus int, err error) {
	if err := s.validateParameter(pagination); err != nil {
		return nil, http.StatusBadRequest, err
	}

	movieProto := &moviepb.GetMovieReq{
		Movie: &moviepb.MovieReq{
			Pagination: pagination,
			Searchword: searchword,
		},
	}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer conn.Close()

	client := moviepb.NewMovieServiceClient(conn)
	result, err = client.GetMovie(context.Background(), movieProto)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return result, http.StatusOK, nil
}

func (s ServiceImpl) validateParameter(pagination string) (err error) {
	page, err := strconv.Atoi(pagination)
	if err != nil {
		return errors.New("pagination must be integer")
	}

	if page < 1 || page > 100 {
		return errors.New("pagination must be 1-100")
	}
	return
}
