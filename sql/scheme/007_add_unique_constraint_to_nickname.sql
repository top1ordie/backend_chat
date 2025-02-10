-- +goose Up 

ALTER TABLE users ADD CONSTRAINT uniq_name UNIQUE(nickname);

-- +goose Down 
ALTER TABLE users DELETE CONSTRAINT uniq_name;
