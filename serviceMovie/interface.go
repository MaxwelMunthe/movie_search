package serviceMovie

import (
	"context"
	"net/http"

	"movie-search/proto"
)

type Service interface {
	GetMovie(ctx context.Context, req *moviepb.GetMovieReq) (*moviepb.GetMovieRes, error)
}

type Repository interface {
	GetDataFromOmdbAPI(pagination, searchWord string) (resp *http.Response, err error)
}
