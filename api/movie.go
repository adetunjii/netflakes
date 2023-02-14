package api

import (
	"errors"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/adetunjii/netflakes/model"
	"github.com/gin-gonic/gin"
)

type ByReleaseDate []model.Movie

func (m ByReleaseDate) Len() int      { return len(m) }
func (m ByReleaseDate) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m ByReleaseDate) Less(i, j int) bool {

	iReleaseDate, err := time.Parse("2023-01-01", m[i].ReleaseDate)
	if err != nil {
		return false
	}

	jReleaseDate, err := time.Parse("2023-00-01", m[j].ReleaseDate)
	if err != nil {
		return false
	}

	return iReleaseDate.Before(jReleaseDate)
}

type ByReverseChronologicalOrder []model.Comment

func (c ByReverseChronologicalOrder) Len() int      { return len(c) }
func (c ByReverseChronologicalOrder) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByReverseChronologicalOrder) Less(i, j int) bool {
	return c[j].CreatedAt.Before(c[i].CreatedAt)
}

func (s *Server) fetchMovies(ctx *gin.Context) {

	movies, err := s.kvstore.GetMovies(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else if len(movies) == 0 {
		movies, err = s.movieApi.FetchMovies(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		sortedMovies := []model.Movie{}
		for _, movie := range movies {
			count, err := s.sqlstore.Movie().FetchCommentCounts(ctx, int64(movie.EpisodeID))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			movie.CommentCount = count
			sortedMovies = append(sortedMovies, movie)
		}

		// store the movies to cache for subsequent request
		if err := s.kvstore.SetMovies(ctx, sortedMovies); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		sort.Sort(ByReleaseDate(sortedMovies))
		ctx.JSON(http.StatusOK, sortedMovies)
		return
	}

	sort.Sort(ByReleaseDate(movies))
	ctx.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}

type addCommentParams struct {
	MovieID int64 `uri:"movie_id" binding:"required,min=1"`
}

type addCommentRequest struct {
	Body      string `json:"body" binding:"required,max=500"`
	CreatedBy string `json:"created_by" binding:"required"`
}

func (s *Server) addComment(ctx *gin.Context) {
	var req addCommentRequest
	var uri addCommentParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// assumes the episode_id of the movie is the movie_id
	movie, err := s.movieApi.FetchMovie(ctx, uri.MovieID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment := &model.Comment{
		MovieID:   int64(movie.EpisodeID),
		MovieUrl:  movie.Url,
		Body:      req.Body,
		SenderIP:  ctx.ClientIP(),
		CreatedBy: req.CreatedBy,
	}

	if err := s.sqlstore.Movie().SaveComment(ctx, comment); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "comment added succesfully",
	})

}

type getCommentsParams struct {
	MovieID int64 `uri:"movie_id" binding:"required"`
}

type getCommentsQuery struct {
	Page int64 `form:"page" binding:"required,min=1"`
	Size int64 `form:"size" binding:"required,max=100"`
}

func (s *Server) getComments(ctx *gin.Context) {
	var params getCommentsParams
	var query getCommentsQuery

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comments, err := s.sqlstore.Movie().FetchComments(ctx, params.MovieID, query.Page, query.Size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	commentCount, err := s.sqlstore.Movie().FetchCommentCounts(ctx, params.MovieID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// sort the comments in chronological order
	sort.Sort(ByReverseChronologicalOrder(comments))

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "successfully fetched comments",
		"data":          comments,
		"page":          query.Page,
		"size":          query.Size,
		"totalElements": commentCount,
	})

}

type getCharactersParams struct {
	MovieID int64 `uri:"movie_id" binding:"required"`
}

type filterQuery struct {
	SortBy         string `form:"sort_by"`
	Order          string `form:"order"`
	FilterByGender string `form:"filter_by_gender"`
}

func (s *Server) getMovieCharacters(ctx *gin.Context) {
	var params getCharactersParams
	var query filterQuery

	var characters []model.Character
	var err error

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	characters, err = s.kvstore.GetMovieCharacters(ctx, params.MovieID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if len(characters) == 0 {
		characters, err = s.movieApi.FetchMovieCharacters(ctx, params.MovieID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		// save in cache for subsequent requests
		if err := s.kvstore.SetMovieCharacters(ctx, params.MovieID, characters); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	}

	if (query != filterQuery{}) {
		switch query.Order {
		case "asc":
			switch query.SortBy {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name < characters[j].Name
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender < characters[j].Gender
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					xHeight, _ := strconv.Atoi(characters[i].Height)
					yHeight, _ := strconv.Atoi(characters[j].Height)

					return xHeight < yHeight

				})
			default:
				ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid filter query, please use one of name, gender or height")))
				return
			}
		case "desc":
			switch query.SortBy {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[j].Name < characters[i].Name
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[j].Gender < characters[i].Gender
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					xHeight, _ := strconv.Atoi(characters[i].Height)
					yHeight, _ := strconv.Atoi(characters[j].Height)

					return yHeight < xHeight
				})
			default:
				ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid sort query, please use one of name, gender or height")))
				return
			}
		default:
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid order query, please use one of asc or desc")))
			return

		}

		if query.FilterByGender != "" {
			filteredCharacters := []model.Character{}
			totalHeight := 0

			for _, character := range characters {
				if character.Gender == query.FilterByGender {
					filteredCharacters = append(filteredCharacters, character)
					height, _ := strconv.Atoi(character.Height)
					totalHeight += height
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "successfully fetched movie characters",
				"data":    filteredCharacters,
				"metadata": gin.H{
					"totalMatches":   len(filteredCharacters),
					"totalHeight_cm": totalHeight, // assumes the original height is in cm
					"totalHeight_ft": float32(totalHeight) / 30.48,
				},
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully fetched movie characters",
		"data":    characters,
	})
}
