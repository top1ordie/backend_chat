-- +goose Up

CREATE TABLE users (
  id INT PRIMARY KEY,
  nickname TEXT NOT NULL,
  password TEXT NOT NULL
);

-- +goose Down --

DROP TABLE user;
