package serviceMovie

import (
	"net/http"
	"net/url"
	"time"

	"movie-search/infrastructure"
)

type RepositoryImpl struct {
	Config infrastructure.OmdbApi
}

func NewRepository(config infrastructure.OmdbApi) *RepositoryImpl {
	return &RepositoryImpl{Config: config}
}

func (r RepositoryImpl) GetDataFromOmdbAPI(pagination, searchWord string) (resp *http.Response, err error) {
	client := &http.Client{Timeout: time.Second * 30}

	urlParse, err := url.Parse(r.Config.Url)
	if err != nil {
		return nil, err
	}

	query := urlParse.Query()
	query.Add(ParameterApiKey, r.Config.Key)
	query.Add(ParameterSearchWord, searchWord)
	query.Add(ParameterPage, pagination)

	urlParse.RawQuery = query.Encode()
	resp, err = client.Get(urlParse.String())
	return
}
