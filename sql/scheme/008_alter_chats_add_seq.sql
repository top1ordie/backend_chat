-- +goose Up 
CREATE SEQUENCE chats_id_seq START 1;
ALTER TABLE chats ALTER COLUMN id SET DEFAULT nextval('chats_id_seq');
-- +goose Down
DROP SEQUENCE chats_id_seq;
ALTER TABLE chats ALTER COLUMN id DROP DEFAULT;
