-- name: CreateChat :one
INSERT INTO chats(chat_name) VALUES($1)
RETURNING *;
