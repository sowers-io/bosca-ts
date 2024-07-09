-- +goose Up
-- +goose StatementBegin
create table models
(
    id            uuid    not null default gen_random_uuid(),
    type          varchar not null,
    name          varchar not null,
    description   varchar not null,
    configuration jsonb   not null,
    primary key (id)
);

create table prompts
(
    id          uuid    not null default gen_random_uuid(),
    name        varchar not null,
    description varchar not null,
    prompt      text    not null,
    primary key (id)
);

create type storage_system_type as enum ('vector', 'search', 'supplementary');

create table storage_systems
(
    id            uuid                not null default gen_random_uuid(),
    name          varchar             not null,
    description   varchar             not null,
    type          storage_system_type not null,
    configuration jsonb               not null,
    primary key (id)
);

create table storage_system_models
(
    system_id     uuid  not null,
    model_id      uuid  not null,
    configuration jsonb not null default '{
      "type": "default"
    }'::jsonb,
    primary key (system_id, model_id),
    foreign key (system_id) references storage_systems (id),
    foreign key (model_id) references models (id)
);

create table activities
(
    id                varchar not null,
    name              varchar not null,
    description       varchar not null,
    child_workflow_id varchar,
    configuration     jsonb   not null default '{}',
    primary key (id)
);

create type activity_parameter_type as enum ('context', 'supplementary', 'supplementary_array');

create table activity_inputs
(
    activity_id varchar                 not null,
    name        varchar                 not null,
    type        activity_parameter_type not null,
    primary key (activity_id, name),
    foreign key (activity_id) references activities (id) on delete cascade
);

create table activity_outputs
(
    activity_id varchar                 not null,
    name        varchar                 not null,
    type        activity_parameter_type not null,
    primary key (activity_id, name),
    foreign key (activity_id) references activities (id) on delete cascade
);

create table workflows
(
    id            varchar not null, -- This is the identifier of the temporal workflow
    name          varchar not null,
    description   varchar not null,
    queue         varchar not null,
    configuration jsonb   not null default '{}',
    primary key (id)
);

alter table activities
    add foreign key (child_workflow_id) references workflows (id);

create table workflow_activities
(
    id              bigserial not null,
    workflow_id     varchar   not null,
    activity_id     varchar   not null,
    execution_group int       not null,
    configuration   jsonb     not null default '{}',
    primary key (id),
    foreign key (workflow_id) references workflows (id),
    foreign key (activity_id) references activities (id)
);

create table workflow_activity_inputs
(
    activity_id bigint  not null,
    name        varchar not null,
    value       jsonb   not null,
    primary key (activity_id, name),
    foreign key (activity_id) references workflow_activities
);

create table workflow_activity_outputs
(
    activity_id bigint  not null,
    name        varchar not null,
    value       jsonb   not null,
    primary key (activity_id, name),
    foreign key (activity_id) references workflow_activities
);

create index ix_workflow_activities_ix on workflow_activities (workflow_id);

create table workflow_activity_storage_systems
(
    activity_id       bigint not null,
    storage_system_id uuid   not null,
    configuration     jsonb  not null default '{}'::jsonb,
    primary key (activity_id, storage_system_id),
    foreign key (activity_id) references workflow_activities (id),
    foreign key (storage_system_id) references storage_systems (id)
);

create table workflow_activity_models
(
    activity_id   bigint not null,
    model_id      uuid   not null,
    configuration jsonb  not null default '{}'::jsonb,
    primary key (activity_id, model_id),
    foreign key (activity_id) references workflow_activities (id),
    foreign key (model_id) references models (id)
);

create table workflow_activity_prompts
(
    activity_id   bigint not null,
    prompt_id     uuid   not null,
    configuration jsonb  not null default '{}'::jsonb,
    primary key (activity_id, prompt_id),
    foreign key (activity_id) references workflow_activities (id),
    foreign key (prompt_id) references prompts (id)
);

create type workflow_state_type as enum ('processing', 'draft', 'pending', 'approval', 'approved', 'published', 'failure');

create table workflow_states
(
    id                varchar             not null,
    name              varchar             not null,
    description       varchar             not null,
    type              workflow_state_type not null,
    configuration     jsonb               not null default '{}',
    workflow_id       varchar,
    exit_workflow_id  varchar, -- workflow that must return true before exiting
    entry_workflow_id varchar, -- workflow that must return true before entering
    primary key (id),
    foreign key (workflow_id) references workflows (id),
    foreign key (exit_workflow_id) references workflows (id)
);

create table workflow_state_transitions
(
    from_state_id varchar not null,
    to_state_id   varchar not null,
    description   varchar not null,
    primary key (from_state_id, to_state_id),
    foreign key (from_state_id) references workflow_states (id),
    foreign key (to_state_id) references workflow_states (id)
);

create table workflow_executions
(
    id          uuid               default gen_random_uuid(),
    created     timestamp not null default now(),
    workflow_id varchar   not null,
    metadata_id uuid      not null,
    context     jsonb     not null,
    completed   timestamp,
    failed      timestamp,
    error       varchar,
    primary key (id),
    foreign key (workflow_id) references workflows (id)
);

create table workflow_workers
(
    id         uuid               default gen_random_uuid(),
    registered timestamp not null default now(),
    primary key (id)
);

create table workflow_execution_jobs
(
    id                    uuid               default gen_random_uuid(),
    workflow_execution_id uuid      not null,
    workflow_activity_id  bigint    not null,
    activity_id           varchar   not null,
    queue                 varchar   not null,
    created               timestamp not null default now(),
    started               timestamp,
    failed                timestamp,
    completed             timestamp,
    scheduled             timestamp not null,
    tries                 int                default 0,
    worker_id             uuid,
    context               jsonb,
    error                 varchar,
    primary key (id),
    foreign key (workflow_execution_id) references workflow_executions (id),
    foreign key (worker_id) references workflow_workers (id) on delete cascade,
    foreign key (workflow_activity_id) references workflow_activities (id),
    foreign key (activity_id) references activities (id)
);

create index workflow_execution_job_ix on workflow_execution_jobs (queue, activity_id, scheduled asc) where scheduled is null;

create table workflow_execution_job_history
(
    id      uuid               default gen_random_uuid(),
    job_id  uuid      not null,
    created timestamp not null default now(),
    message varchar   not null,
    primary key (id),
    foreign key (job_id) references workflow_execution_jobs (id)
);

create or replace function listen(queue text) returns void AS
$$
begin
    execute format('LISTEN %I;', queue);
end;
$$ language plpgsql;

create or replace function listen_all() returns void AS
$$
declare
    queue varchar;
begin
    for queue in select queue from workflows
        loop
            execute format('LISTEN %I;', queue);
        end loop;
end;
$$ language plpgsql;

create or replace function unlisten(queue text) returns void AS
$$
begin
    execute format('UNLISTEN %I;', queue);
end;
$$ language plpgsql;

create or replace function unlisten_all() returns void AS
$$
declare
    queue varchar;
begin
    for queue in select queue from workflows
        loop
            execute format('UNLISTEN %I;', queue);
        end loop;
end;
$$ language plpgsql;

create or replace function claim_next_job(in_worker_id uuid, in_queue varchar, in_activity_ids varchar[]) returns workflow_execution_jobs as
$$
declare
    job workflow_execution_jobs;
begin

    update workflow_execution_jobs
    set worker_id = in_worker_id,
        started   = now(),
        tries     = tries + 1
    where id = (select id
                from workflow_execution_jobs
                where queue = in_queue
                  and activity_id = any (in_activity_ids)
                order by scheduled for update skip locked
                limit 1)
    returning * into job;

    if job is not null then
        insert into workflow_execution_job_history (job_id, message)
        values (job.id, 'assigned to worker ' || in_worker_id);
    end if;

    return job;
end;
$$ language plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists storage_systems cascade;
drop table if exists storage_system_models cascade;
drop type storage_system_type cascade;
drop table prompts cascade;
drop table if exists workflow_states cascade;
drop type if exists workflow_state_type cascade;
drop table if exists models cascade;
drop table if exists workflows cascade;
drop table if exists workflow_state_transitions cascade;
drop table activities cascade;
drop table workflow_activities cascade;
drop type activity_parameter_type cascade;
drop table activity_inputs cascade;
drop table activity_outputs cascade;
drop table workflow_activity_inputs cascade;
drop table workflow_activity_outputs cascade;
drop table workflow_activity_storage_systems cascade;
drop table workflow_activity_prompts cascade;
drop table workflow_activity_models cascade;
drop table workflow_executions cascade;
drop table workflow_execution_jobs cascade;
drop table workflow_execution_job_history cascade;
drop table workflow_workers cascade;
drop function if exists claim_next_job(in_worker_id uuid, in_queue varchar, in_activity_ids varchar[]);
drop function if exists listen_all();
drop function if exists listen(queue text);
drop function if exists unlisten_all();
drop function if exists unlisten(queue text);
-- +goose StatementEnd
