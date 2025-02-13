-- name: GetUsersInChatById :many
select users.id, users.nickname from users
join user_chat on user_chat.user_id = users.id
where user_chat.chat_id = $1;
-- name: GetUserInChat :one 
select users.id, users.nickname, chats.chat_name from users
join user_chat on users.id = user_chat.user_id
join chats on chats.id = user_chat.chat_id
where chats.id = $1 and users.id = $2;

