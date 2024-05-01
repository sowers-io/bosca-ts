-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgmq CASCADE;
SELECT pgmq.create('metadata');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT pgmq.drop_queue('metadata');
DROP EXTENSION pgmq;
-- +goose StatementEnd
