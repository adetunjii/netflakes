package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adetunjii/netflakes/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestServer_FetchMovies(t *testing.T) {

	testCases := []struct {
		name          string
		buildStubs    func(kvstore *mock.MockKVStore, swapi *mock.MockMovieApi)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Status OK",
			buildStubs: func(kvstore *mock.MockKVStore, swapi *mock.MockMovieApi) {
				kvstore.EXPECT().GetMovies(gomock.Any()).Times(1)
				swapi.EXPECT().FetchMovies(gomock.Any()).Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			swapi := mock.NewMockMovieApi(ctrl)
			kvstore := mock.NewMockKVStore(ctrl)
			sqlstore := mock.NewMockSqlStore(ctrl)

			testCase.buildStubs(kvstore, swapi)

			server := NewTestServer(t, kvstore, sqlstore, swapi)
			recorder := httptest.NewRecorder()

			url := "/movies"

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)

		})
	}

}
