package api

import (
	"os"
	"testing"

	"github.com/adetunjii/netflakes/pkg/logging"
	"github.com/adetunjii/netflakes/port"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, kvstore port.KVStore, sqlstore port.SqlStore, swapi port.MovieApi) *Server {

	zapSugarLogger := logging.NewZapSugarLogger()
	logger := logging.NewLogger(zapSugarLogger)
	require.NotEmpty(t, logger)

	server := NewServer(kvstore, sqlstore, swapi, logger)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
