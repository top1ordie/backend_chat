-- name: GetUsersChats :many
SELECT * FROM chats WHERE id IN (SELECT chat_id FROM user_chat WHERE user_id = $1);
