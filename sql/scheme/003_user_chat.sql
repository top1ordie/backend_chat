-- +goose Up 
CREATE TABLE user_chat (
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  chat_id INT REFERENCES chats(id) ON DELETE CASCADE,
  PRIMARY KEY(user_id, chat_id)
);
-- +goose Down
DROP TABLE user_chat;
