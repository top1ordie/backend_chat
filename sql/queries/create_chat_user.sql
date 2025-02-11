-- name: CreateUser_Chat :one
INSERT INTO user_chat (
  user_id,
  chat_id
) VALUES ( $1,$2 )
RETURNING *;
