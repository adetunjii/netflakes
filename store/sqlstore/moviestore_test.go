package sqlstore

import (
	"context"
	"testing"

	"github.com/adetunjii/netflakes/model"
	"github.com/stretchr/testify/require"
)

var testComment = model.Comment{
	ID:       int64(1),
	Body:     "hello guys, this is just a test comment. Absolutely love it!!",
	MovieID:  int64(1),
	MovieUrl: "https://swapi.dev/api/films/1",
	SenderIP: "192.168.0.0",
}

func TestSaveComment(t *testing.T) {

	err := sqlStore.Movie().SaveComment(context.Background(), &testComment)
	require.NoError(t, err)
}

func TestFetchComment(t *testing.T) {
	comments, err := sqlStore.Movie().FetchComments(context.Background(), testComment.MovieID, 1, 20)
	require.NoError(t, err)
	require.NotEmpty(t, comments)
	require.Len(t, comments, 1)

	comment := comments[0]
	require.NotEmpty(t, comment.ID)
	require.NotZero(t, comment.ID)
	require.Equal(t, testComment.MovieID, comment.MovieID)
	require.Equal(t, testComment.MovieUrl, comment.MovieUrl)
	require.Equal(t, testComment.SenderIP, comment.SenderIP)
}

func TestGetComment(t *testing.T) {
	comment, err := sqlStore.Movie().GetComment(context.Background(), testComment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, testComment.ID, comment.ID)
	require.Equal(t, testComment.MovieUrl, comment.MovieUrl)
	require.Equal(t, testComment.MovieID, comment.MovieID)
	require.Equal(t, testComment.SenderIP, comment.SenderIP)

	// fetch non-existent comment

	_, err = sqlStore.Movie().GetComment(context.Background(), 4)
	require.Error(t, err)
}

func TestDeleteComment(t *testing.T) {
	err := sqlStore.Movie().DeleteComment(context.Background(), testComment.ID)
	require.NoError(t, err)

	comment, err := sqlStore.Movie().GetComment(context.Background(), testComment.ID)
	require.Error(t, err)
	require.Empty(t, comment)
}
