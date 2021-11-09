-- name: GetQuote :one
SELECT * FROM quotes
WHERE id = $1 LIMIT 1;

-- name: GetQuoteRows :one
SELECT COUNT(*)
FROM quotes;