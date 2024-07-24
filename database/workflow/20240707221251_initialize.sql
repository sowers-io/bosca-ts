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
    id            uuid    not null default gen_random_uuid(),
    name          varchar not null,
    description   varchar not null,
    system_prompt text    not null,
    user_prompt   text    not null,
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
    queue         varchar not null check (length(queue) > 0),
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
    queue           varchar,
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
    value       varchar not null,
    primary key (activity_id, name),
    foreign key (activity_id) references workflow_activities
);

create table workflow_activity_outputs
(
    activity_id bigint  not null,
    name        varchar not null,
    value       varchar not null,
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
-- +goose StatementEnd
