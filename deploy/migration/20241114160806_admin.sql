-- +goose Up
-- +goose StatementBegin

create TABLE IF NOT EXISTS admins
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR,
    email      VARCHAR, -- after end work do them unique
    phone      VARCHAR,
    hashed_password VARCHAR,
    photo      VARCHAR
);


create TABLE IF NOT EXISTS admin_hotel
(
    id_admin int,
    id_hotel int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists admins, admin_hotel;
-- +goose StatementEnd
