-- +goose Up
CREATE TABLE text_message (
  ID SERIAL PRIMARY KEY,
  MESSAGE_ID INT REFERENCES messages(id),
  DATA TEXT
);
CREATE TABLE image_message (
  ID SERIAL PRIMARY KEY,
  MESSAGE_ID INT REFERENCES messages(id),
  DATA_URL TEXT
);
CREATE TABLE media_message (
  ID SERIAL PRIMARY KEY,
  MESSAGE_ID INT REFERENCES messages(id),
  DATA_URL TEXT
);

-- +goose Down 
DROP TABLE text_message;
DROP TABLE image_message;
DROP TABLE media_message;
