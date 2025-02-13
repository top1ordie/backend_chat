-- name: CreateMessage :one
INSERT INTO messages(chat_id,user_id,created_at) VALUES($1,$2,$3)
RETURNING messages.id;

-- name: CreateTextMessage :one
INSERT INTO text_message(message_id,data) VALUES($1 , $2)
RETURNING *;
-- name: CreateImageMessage :one
INSERT INTO image_message(message_id,data_url) VALUES($1 , $2)
RETURNING *;
-- name: CreateMediaMessage :one
INSERT INTO media_message(message_id,data_url) VALUES($1 , $2)
RETURNING *;
