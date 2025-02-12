-- +goose Up 
ALTER TABLE user_chat DROP CONSTRAINT user_chat_pkey;
-- +goose Down 
ALTER TABLE user_chat ADD CONSTRAINT user_chat_pkey PRIMARY KEY(user_id,chat_id);
