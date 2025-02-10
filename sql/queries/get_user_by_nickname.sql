-- name: GetUserByNickName :one
SELECT * FROM users WHERE nickname = $1;
