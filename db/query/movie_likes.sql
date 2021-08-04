-- name: CreateMovieLikes :one
INSERT INTO movie_likes (
  movie_id,
  user_name,
  liked
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetMovieForUser :one
SELECT * FROM movie_likes
WHERE user_name = $1 and movie_id = $2;
