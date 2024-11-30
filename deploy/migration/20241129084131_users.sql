-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    email VARCHAR,
    hashed_pwd VARCHAR,
    phone VARCHAR,
    photo VARCHAR
);

CREATE TABLE IF NOT EXISTS users_rooms(
    id_user BIGINT,
    id_room BIGINT,
    day_checkin DATE,
    day_checkout DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists users_rooms;
-- +goose StatementEnd
