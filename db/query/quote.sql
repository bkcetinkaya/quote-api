-- name: GetQuote :one
SELECT * FROM quotes
WHERE id = $1 LIMIT 1;