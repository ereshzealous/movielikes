-- name: CreateMovie :one
INSERT INTO movie (
  name,
  production_company,
  year_released
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetMovie :one
SELECT * FROM movie WHERE id = $1;

-- name: GetMoviesFromProductionCompany :many
SELECT * FROM movie WHERE production_company = $1 ORDER BY year_released DESC limit $2 offset $3;

-- name: GetMoviesByReleasedYear :many
SELECT * FROM movie WHERE year_released = $1 ORDER BY created_at DESC limit $2 offset $3;