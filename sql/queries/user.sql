-- name: CreateUser :one
INSERT INTO users (id,created_at,updated_at, name,api_key)
VALUES($1,$2,$3,$4,$5)
RETURNING *;

