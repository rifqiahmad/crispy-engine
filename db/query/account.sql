-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING id, owner, balance, currency, created_at;

-- name: GetAccount :one
SELECT
    id,
    owner,
    balance,
    currency,
    created_at
FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT
    id,
    owner,
    balance,
    currency,
    created_at
FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT
    id,
    owner,
    balance,
    currency,
    created_at
FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
    SET balance = $2
    WHERE id = $1
RETURNING id, owner, balance, currency, created_at;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING id, owner, balance, currency, created_at;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
