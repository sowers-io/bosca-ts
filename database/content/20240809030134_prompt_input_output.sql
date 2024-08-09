-- +goose Up
-- +goose StatementBegin
alter table prompts add column
    input_type varchar;
alter table prompts add column
    output_type varchar;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table prompts drop column input_type;
alter table prompts drop column output_type;
-- +goose StatementEnd
