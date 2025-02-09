-- name: CreateUser :one
INSERT INTO users (id,nickname,password) 
VALUES ($1,$2,$3)
RETURNING *;
