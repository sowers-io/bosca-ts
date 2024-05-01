-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgmq CASCADE;
SELECT pgmq.create('metadata');
CREATE OR REPLACE FUNCTION listen(channel text) RETURNS void AS $$
BEGIN
    EXECUTE format('LISTEN %I;', channel);
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION unlisten(channel text) RETURNS void AS $$
BEGIN
    EXECUTE format('UNLISTEN %I;', channel);
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT pgmq.drop_queue('metadata');
DROP EXTENSION pgmq;
DROP FUNCTION IF EXISTS listen(channel text);
DROP FUNCTION IF EXISTS unlisten(channel text);
-- +goose StatementEnd
