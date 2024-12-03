-- +goose Up
-- +goose StatementBegin

create TABLE IF NOT EXISTS hotels
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR,
    stars   int,
    rating  float default 0,
    address VARCHAR,
    email   VARCHAR,
    phone   VARCHAR,
    links   JSONB,
    lat     varchar,
    lon     varchar
);


create TABLE IF NOT EXISTS photo_hotels
(
    id       SERIAL PRIMARY KEY,
    hotel_id int,
    name     VARCHAR,
    photo    varchar
);
create TABLE IF NOT EXISTS ratings
(
    id       serial primary key,
    id_user  int,
    id_hotel int,
    rating   float
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists hotels, photo_hotels, ratings;
-- +goose StatementEnd
