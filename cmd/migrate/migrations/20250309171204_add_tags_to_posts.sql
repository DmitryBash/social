-- +goose Up
-- +goose StatementBegin
ALTER TABLE
  posts
ADD
  COLUMN tags VARCHAR(100) [];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts DROP COLUMN tags; 
-- +goose StatementEnd
