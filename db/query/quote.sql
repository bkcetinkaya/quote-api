-- name: GetQuote :one
SELECT * FROM quotes
WHERE id = $1 LIMIT 1;

-- name: GetQuoteRows :one
SELECT COUNT(*)
FROM quotes;

-- name: CreateQuote :one
INSERT INTO quotes (
  quote, author
) VALUES (
  $1, $2
)
RETURNING *;