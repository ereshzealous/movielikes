package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomMovieLike(t *testing.T) MovieLike {
	movie := CreateRandomMovie(t);
	user := CreateRandomUser(t);

	arg := CreateMovieLikesParams {
		MovieID: movie.ID,
		UserName: user.UserName,
		Liked: true,
	}
	movieLike, err := testQueries.CreateMovieLikes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movieLike)
	require.Equal(t, arg.MovieID, movieLike.MovieID)
	require.Equal(t, arg.UserName, movieLike.UserName)
	require.Equal(t, arg.Liked, movieLike.Liked)
	require.NotZero(t, movieLike.ID)
	require.NotZero(t, movieLike.CreatedAt)
	return movieLike;
}

func TestCreateMovieLikes(t *testing.T) {
	CreateRandomMovieLike(t)
}

func TestGetMovieLikeForUser(t *testing.T)  {
	movieLike := CreateRandomMovieLike(t)
	arg := GetMovieForUserParams{
		UserName: movieLike.UserName,
		MovieID: movieLike.MovieID,
	}
	like, err := testQueries.GetMovieForUser(context.Background(), arg);
	require.NoError(t, err)
	require.NotEmpty(t, like)
	require.Equal(t, movieLike.MovieID, like.MovieID)
	require.Equal(t, movieLike.UserName, like.UserName)
	require.Equal(t, movieLike.Liked, like.Liked)
	require.Equal(t, movieLike.ID, like.ID)
	require.Equal(t, movieLike.CreatedAt, like.CreatedAt)
}