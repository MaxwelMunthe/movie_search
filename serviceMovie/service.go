package serviceMovie

import (
	"context"
	"encoding/json"

	"movie-search/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type MovieServiceServer struct {
	Repository Repository
	Logger log.Logger
}

func NewService(repository Repository) *MovieServiceServer {
	return &MovieServiceServer{Repository: repository}
}

func (s *MovieServiceServer) GetMovie(ctx context.Context, req *moviepb.GetMovieReq) (response *moviepb.GetMovieRes, err error) {
	logger := log.With(s.Logger, "method", "Get")
	result, err := s.Repository.GetDataFromOmdbAPI(req.Movie.Pagination, req.Movie.Searchword)
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err = json.NewDecoder(result.Body).Decode(&response); err != nil {
		level.Error(logger).Log("err", err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return response, nil
}
