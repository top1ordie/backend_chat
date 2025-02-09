-- +goose Up
CREATE SEQUENCE user_id_seq START 2;
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('user_id_seq');
-- +goose Down
DROP SEQUENCE user_id_seq;
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
