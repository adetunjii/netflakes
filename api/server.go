package api

import (
	"github.com/adetunjii/netflakes/port"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	kvstore  port.KVStore
	sqlstore port.SqlStore
	movieApi port.MovieApi
	logger   port.Logger
}

func NewServer(kvstore port.KVStore, sqlstore port.SqlStore, movieApi port.MovieApi, logger port.Logger) *Server {
	server := &Server{
		kvstore:  kvstore,
		sqlstore: sqlstore,
		movieApi: movieApi,
		logger:   logger,
	}

	server.setupRouter()
	return server
}

func (s *Server) Start() error {
	return s.router.Run()
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.GET("/api/movies", s.fetchMovies)
	router.POST("/api/movies/:movie_id/add-comment", s.addComment)
	router.GET("/api/movies/:movie_id/comments", s.getComments)
	router.GET("/api/movies/:movie_id/characters", s.getMovieCharacters)

	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
