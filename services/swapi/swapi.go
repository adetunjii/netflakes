package swapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/adetunjii/netflakes/port"
)

type Swapi struct {
	baseURL string
	client  *http.Client
	logger  port.Logger
}

var _ port.MovieApi = (*Swapi)(nil)

func New(baseUrl string, logger port.Logger) *Swapi {
	return &Swapi{
		baseURL: baseUrl,
		client:  &http.Client{},
		logger:  logger,
	}
}

func (s *Swapi) Get(path string) ([]byte, error) {

	route := path
	var err error
	// check if url contains https://
	if path[:5] != "https" {
		parsedRoute, err := url.Parse(s.baseURL + path)
		if err != nil {
			return nil, err
		}
		route = parsedRoute.String()
	}

	request, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, errors.New("resource not found")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
