-- +goose Up 
CREATE TABLE messages (
  id INT PRIMARY KEY,
  chat_id INT REFERENCES chats(id),
  user_id INT REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE messages;
