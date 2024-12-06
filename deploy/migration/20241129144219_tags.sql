-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags(
    id   SERIAL PRIMARY KEY,
    id_type int,
    type VARCHAR,
    tag VARCHAR
);

Create TABLE if not exists tag_type(
    id serial primary key,
    type varchar
);

CREATE TABLE IF NOT EXISTS hotels_tags(
    id_hotel BIGINT,
    id_tag   BIGINT
);

CREATE TABLE IF NOT EXISTS rooms_tags(
    id_room BIGINT,
    id_tag   BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tags;
drop table if exists hotels_tags;
drop table if exists rooms_tags;
-- +goose StatementEnd
