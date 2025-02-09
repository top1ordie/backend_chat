-- +goose Up 
CREATE TABLE chats (
  id INT PRIMARY KEY,
  chat_name TEXT
);

-- +goose Down
DROP TABLE chats;
