-- +goose Up 
CREATE SEQUENCE messages_id_seq START 1;
ALTER TABLE messages ALTER COLUMN id SET DEFAULT nextval('messages_id_seq');
-- +goose Down
DROP SEQUENCE messages_id_seq;
ALTER TABLE messages ALTER COLUMN id DROP DEFAULT;
