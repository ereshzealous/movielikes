package db

import (
	"context"
	"testing"

	"example.com/movielikes/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomMovie(t *testing.T) Movie {
	arg := CreateMovieParams{
		Name: util.RandomString(10),
		ProductionCompany: util.RandomString(12),
		YearReleased: util.RandomYear(),
	}
	movie, err := testQueries.CreateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movie)
	require.Equal(t, arg.Name, movie.Name)
	require.Equal(t, arg.ProductionCompany, movie.ProductionCompany)
	require.Equal(t, arg.YearReleased, movie.YearReleased)
	require.NotZero(t, movie.ID)
	require.NotZero(t, movie.CreatedAt)
	return movie
}

func createNewMovie(t *testing.T, YearReleased int32, ProductionCompany string) Movie {
	arg := CreateMovieParams{
		Name: util.RandomString(10),
		ProductionCompany: ProductionCompany,
		YearReleased: YearReleased,
	}
	movie, err := testQueries.CreateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movie)
	require.Equal(t, arg.Name, movie.Name)
	require.Equal(t, arg.ProductionCompany, movie.ProductionCompany)
	require.Equal(t, arg.YearReleased, movie.YearReleased)
	require.NotZero(t, movie.ID)
	require.NotZero(t, movie.CreatedAt)
	return movie
}
func TestCreateMovie(t *testing.T) {
	CreateRandomMovie(t);
}

func TestGetMovie(t *testing.T) {
	expectedMovie := CreateRandomMovie(t)
	actualMovie, err := testQueries.GetMovie(context.Background(), expectedMovie.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualMovie)
	require.Equal(t, expectedMovie.Name, actualMovie.Name)
	require.Equal(t, expectedMovie.ProductionCompany, actualMovie.ProductionCompany)
	require.Equal(t, expectedMovie.YearReleased, actualMovie.YearReleased)
	require.Equal(t, expectedMovie.ID, actualMovie.ID)
	require.Equal(t, expectedMovie.CreatedAt, actualMovie.CreatedAt)
}

func TestGetMoviesByReleasedYear(t *testing.T) {
	movies := []Movie{}
	for i := 0; i < 5; i++ {
		movies = append(movies, createNewMovie(t, 2020, "TestGetMoviesByReleasedYear"));
	}
	arg := GetMoviesByReleasedYearParams{
		YearReleased: 2020,
		Limit: 2,
		Offset: 0,
	}
	actualMovies, err := testQueries.GetMoviesByReleasedYear(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualMovies)
	require.Len(t, actualMovies, 2)
	for _, movie := range actualMovies {
		require.NotEmpty(t, movie)
		require.Equal(t, movie.YearReleased, int32(2020))
	}
}

func TestGetMoviesFromProductionCompany(t *testing.T) {
	movies := []Movie{}
	for i := 0; i < 5; i++ {
		movies = append(movies, createNewMovie(t, util.RandomYear(), "TestGetMoviesFromProductionCompany"));
	}
	arg := GetMoviesFromProductionCompanyParams{
		ProductionCompany: "TestGetMoviesFromProductionCompany",
		Limit: 2,
		Offset: 0,
	}
	actualMovies, err := testQueries.GetMoviesFromProductionCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualMovies)
	require.Len(t, actualMovies, 2)
	for _, movie := range actualMovies {
		require.NotEmpty(t, movie)
		require.Equal(t, movie.ProductionCompany, "TestGetMoviesFromProductionCompany")
	}
}