-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
