-- name: CreateAccountDetails :one
INSERT INTO account_details (
  user_name,
  balance
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetAccountDetails :one
SELECT * FROM account_details
WHERE id = $1 LIMIT 1;

-- name: GetAccountDetailsForUpdate :one
SELECT * FROM account_details
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccountDetails :many
SELECT * FROM account_details
WHERE user_name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccountDetails :one
UPDATE account_details
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountBalance :one
UPDATE account_details
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccountDetails :exec
DELETE FROM account_details
WHERE id = $1;
