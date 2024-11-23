-- +goose Up
-- +goose StatementBegin
create table widgets (
    id serial primary key,
    color text 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table widgets;
-- +goose StatementEnd
