-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rooms(
    id SERIAL PRIMARY KEY,
    id_hotel BIGINT,
    number INT,
    class VARCHAR,
    room_cnt INT,
    bath_cnt INT,
    bed_cnt INT,
    time_in TIME,
    time_out TIME,
    price FLOAT
);

CREATE TABLE IF NOT EXISTS photo_rooms(
    id SERIAL PRIMARY KEY,
    id_room BIGINT,
    name VARCHAR,
    photo VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists rooms;
drop table if exists photo_rooms;
-- +goose StatementEnd
