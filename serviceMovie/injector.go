package serviceMovie

import (
	"movie-search/infrastructure"
)

func InjectServiceMovie() *MovieServiceServer {
	repository := NewRepository(infrastructure.OmdbAPI)
	service := NewService(repository)

	return service
}
