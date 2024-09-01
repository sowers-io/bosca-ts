-- +goose Up
-- +goose StatementBegin
create table metadata_workflow_jobs (
  id uuid not null,
  job_id varchar not null,
  queue varchar not null,
  created timestamp not null default now(),
  primary key (id, job_id),
  foreign key (id) references metadata(id) on delete cascade
);
create table collection_workflow_jobs (
  id uuid not null,
  job_id varchar not null,
  queue varchar not null,
  created timestamp not null default now(),
  primary key (id, job_id),
  foreign key (id) references collections(id) on delete cascade
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table metadata_workflow_jobs;
drop table collection_workflow_jobs;
-- +goose StatementEnd