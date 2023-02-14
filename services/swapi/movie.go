package swapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adetunjii/netflakes/model"
)

func (s *Swapi) FetchMovies(ctx context.Context) ([]model.Movie, error) {
	path := "/films"
	movies := []model.Movie{}

	response, err := s.Get(path)
	if err != nil {
		return nil, err
	}

	jsonResponse := map[string]interface{}{}
	if err = json.Unmarshal(response, &jsonResponse); err != nil {
		return nil, err
	}

	movieList := jsonResponse["results"]
	j, err := json.Marshal(movieList)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(j, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *Swapi) FetchMovie(ctx context.Context, movieID int64) (*model.Movie, error) {
	path := fmt.Sprintf("/films/%d", movieID)
	movie := &model.Movie{}

	response, err := s.Get(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Swapi) FetchMovieCharacters(ctx context.Context, movieID int64) ([]model.Character, error) {
	path := fmt.Sprintf("/films/%d", movieID)
	movie := &model.Movie{}

	response, err := s.Get(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response, movie); err != nil {
		return nil, err
	}

	characters := []model.Character{}
	for _, characterUrl := range movie.Characters {
		character, err := s.GetCharacter(ctx, characterUrl)
		if err != nil {
			return nil, err
		}
		characters = append(characters, *character)
	}

	return characters, nil
}

func (s *Swapi) GetCharacter(ctx context.Context, url string) (*model.Character, error) {

	character := &model.Character{}
	response, err := s.Get(url)
	if err != nil {
		s.logger.Error("failed to get character: %w", err)
		return nil, err
	}

	if err = json.Unmarshal(response, character); err != nil {
		s.logger.Error("failed to unmarshal json: %w", err)
		return nil, err
	}

	return character, nil

}
