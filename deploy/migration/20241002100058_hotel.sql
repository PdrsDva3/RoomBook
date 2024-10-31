-- +goose Up
-- +goose StatementBegin

create TABLE IF NOT EXISTS hotels
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR,
    stars   int,
    address VARCHAR,
    email   VARCHAR,
    phone   VARCHAR,
    links   JSONB,
    photo   JSONB
);


create TABLE IF NOT EXISTS photos
(
    id       SERIAL PRIMARY KEY,
    list_id  int,
    hotel_id int,
    name     VARCHAR,
    photo    varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists hotels;
drop table if exists photos;
-- +goose StatementEnd
