-- +goose Up

CREATE TABLE products (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  quantity bigint NOT NULL,
  price bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT(now())
);


-- +goose Down
DROP TABLE products;