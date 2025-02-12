-- +goose Up 
ALTER TABLE chats ALTER COLUMN chat_name SET NOT NULL;
-- +goose Down 
ALTER TABLE chats ALTER COLUMN chat_name DROP NOT NULL;
