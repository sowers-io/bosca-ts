-- +goose Up
-- +goose StatementBegin
create table groups
(
    name        varchar,
    description varchar,
    created     timestamp not null default now(),
    enabled     boolean   not null default true,
    primary key (name)
);

insert into groups (name, description) values ('administrators', 'Bosca Administrators');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table groups;
-- +goose StatementEnd
