// Code generated by sqlc. DO NOT EDIT.
// source: movie_likes.sql

package db

import (
	"context"
)

const createMovieLikes = `-- name: CreateMovieLikes :one
INSERT INTO movie_likes (
  movie_id,
  user_name,
  liked
) VALUES (
  $1, $2, $3
) RETURNING id, movie_id, user_name, liked, created_at
`

type CreateMovieLikesParams struct {
	MovieID  int64  `json:"movie_id"`
	UserName string `json:"user_name"`
	Liked    bool   `json:"liked"`
}

func (q *Queries) CreateMovieLikes(ctx context.Context, arg CreateMovieLikesParams) (MovieLike, error) {
	row := q.db.QueryRowContext(ctx, createMovieLikes, arg.MovieID, arg.UserName, arg.Liked)
	var i MovieLike
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.UserName,
		&i.Liked,
		&i.CreatedAt,
	)
	return i, err
}

const getMovieForUser = `-- name: GetMovieForUser :one
SELECT id, movie_id, user_name, liked, created_at FROM movie_likes
WHERE user_name = $1 and movie_id = $2
`

type GetMovieForUserParams struct {
	UserName string `json:"user_name"`
	MovieID  int64  `json:"movie_id"`
}

func (q *Queries) GetMovieForUser(ctx context.Context, arg GetMovieForUserParams) (MovieLike, error) {
	row := q.db.QueryRowContext(ctx, getMovieForUser, arg.UserName, arg.MovieID)
	var i MovieLike
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.UserName,
		&i.Liked,
		&i.CreatedAt,
	)
	return i, err
}
