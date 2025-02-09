-- name: CreateUserNoId :one
INSERT INTO users (
  nickname,password
) VALUES ( $1,$2 )
RETURNING *;
