-- Copyright 2024 Sowers, LLC
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- +goose Up
-- +goose StatementBegin
create table workflows
(
    id            varchar not null, -- This is the identifier of the temporal workflow
    name          varchar not null,
    description   varchar not null,
    queue         varchar not null,
    configuration jsonb   not null default '{}',
    primary key (id)
);

insert into workflows (id, name, description, queue)
values ('workflow.ProcessMetadata', 'Process Metadata', 'Process Metadata', 'workflow');

insert into workflows (id, name, description, queue)
values ('bible.ProcessBible', 'Process Bible', 'Process Bible', 'bible');

create table traits
(
    id          uuid    not null default gen_random_uuid(),
    name        varchar not null,
    workflow_id varchar not null,
    primary key (id),
    foreign key (workflow_id) references workflows (id)
);

insert into traits (name, workflow_id)
values ('USX', 'bible.ProcessBible');

create table categories
(
    id   uuid    not null default gen_random_uuid(),
    name varchar not null,
    primary key (id)
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

insert into workflow_states (id, name, description, type, workflow_id)
values ('processing', 'Processing', 'Initial Processing after Creation', 'processing'::workflow_state_type,
        'workflow.ProcessMetadata');

insert into workflow_states (id, name, description, type)
values ('pending', 'Pending', 'Pending', 'pending'::workflow_state_type),
       ('draft', 'Draft', 'Draft', 'draft'::workflow_state_type),
       ('pending_approval', 'Pending Approval', 'Pending Approval', 'approval'::workflow_state_type),
       ('approved', 'Approved', 'Approved', 'approved'::workflow_state_type),
       ('published', 'Published', 'Published', 'published'::workflow_state_type),
       ('failure', 'Failure', 'Failure', 'failure'::workflow_state_type);

insert into workflow_state_transitions (from_state_id, to_state_id, description)
values ('pending', 'processing', 'content is ready to be processed'),
       ('processing', 'draft', 'content has been processed, now in draft mode'),
       ('processing', 'failure', 'processing failed'),
       ('failure', 'processing', 'processing again'),
       ('draft', 'published', 'publishing draft');

create
    type collection_type as enum ('root', 'standard', 'folder', 'queue');

create table collections
(
    id                        uuid      not null default gen_random_uuid(),
    name                      varchar   not null,
    type                      collection_type    default 'standard',
    attributes                jsonb     not null default '{}',
    tags                      varchar[] not null default '{}',
    created                   timestamp          default now(),
    modified                  timestamp          default now(),
    enabled                   boolean            default true,
    workflow_state_id         varchar   not null default 'pending',
    workflow_state_pending_id varchar,
    primary key (id),
    foreign key (workflow_state_id) references workflow_states (id),
    foreign key (workflow_state_pending_id) references workflow_states (id)
);

create table collection_collection_items
(
    collection_id uuid,
    child_id      uuid,
    primary key (collection_id, child_id),
    foreign key (collection_id) references collections (id) on delete cascade,
    foreign key (child_id) references collections (id) on delete cascade
);

create table collection_traits
(
    collection_id uuid,
    trait_id      uuid,
    primary key (collection_id, trait_id),
    foreign key (collection_id) references collections (id) on delete cascade,
    foreign key (trait_id) references traits (id) on delete cascade
);

create table collection_categories
(
    collection_id uuid,
    category_id   uuid,
    primary key (collection_id, category_id),
    foreign key (collection_id) references collections (id) on delete cascade,
    foreign key (category_id) references categories (id) on delete cascade
);

create
    type metadata_type as enum ('standard', 'variant');

create table metadata
(
    id                        uuid      not null                                  default gen_random_uuid(),
    parent_id                 uuid,
    name                      varchar   not null check (length(name) > 0),
    type                      metadata_type                                       default 'standard',
    content_type              varchar   not null check (length(content_type) > 0),
    content_length            bigint    not null check (content_length > 0),
    language_tag              varchar   not null check (length(language_tag) > 0) default 'en',
    tags                      varchar[] not null                                  default '{}',
    attributes                jsonb     not null                                  default '{}',
    created                   timestamp                                           default now(),
    modified                  timestamp                                           default now(),
    workflow_state_id         varchar   not null                                  default 'pending',
    workflow_state_pending_id varchar,
    source                    varchar,
    primary key (id),
    foreign key (parent_id) references metadata (id) on delete cascade,
    foreign key (workflow_state_id) references workflow_states (id),
    foreign key (workflow_state_pending_id) references workflow_states (id)
);

create table metadata_workflow_transition_history
(
    id            bigserial not null,
    metadata_id   uuid      not null,
    to_state_id   varchar   not null,
    from_state_id varchar   not null,
    subject       varchar   not null,
    status        varchar,
    success       boolean   not null default false,
    complete      boolean   not null default false,
    created       timestamp          default now(),
    primary key (id),
    foreign key (metadata_id) references metadata (id) on delete cascade,
    foreign key (to_state_id) references workflow_states (id),
    foreign key (from_state_id) references workflow_states (id)
);

create table collection_workflow_transition_history
(
    id            bigserial not null,
    metadata_id   uuid      not null,
    to_state_id   varchar   not null,
    from_state_id varchar   not null,
    subject       varchar   not null,
    status        varchar,
    success       boolean   not null default false,
    complete      boolean   not null default false,
    created       timestamp          default now(),
    primary key (id),
    foreign key (metadata_id) references metadata (id) on delete cascade,
    foreign key (to_state_id) references workflow_states (id),
    foreign key (from_state_id) references workflow_states (id)
);

create table collection_metadata_items
(
    collection_id uuid,
    metadata_id   uuid,
    primary key (collection_id, metadata_id),
    foreign key (collection_id) references collections (id) on delete cascade,
    foreign key (metadata_id) references metadata (id) on delete cascade
);

create table metadata_relationship
(
    metadata1_id uuid,
    metadata2_id uuid,
    relationship varchar,
    primary key (metadata1_id, metadata2_id, relationship),
    foreign key (metadata1_id) references metadata (id) on delete cascade,
    foreign key (metadata2_id) references metadata (id) on delete cascade
);

create table metadata_traits
(
    metadata_id uuid,
    trait_id    uuid,
    primary key (metadata_id, trait_id),
    foreign key (metadata_id) references metadata (id) on delete cascade,
    foreign key (trait_id) references traits (id) on delete cascade
);

create table metadata_categories
(
    metadata_id uuid,
    category_id uuid,
    primary key (metadata_id, category_id),
    foreign key (metadata_id) references metadata (id) on delete cascade,
    foreign key (category_id) references categories (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists workflows cascade;
drop table if exists workflow_state_transitions cascade;
drop table if exists collection_traits cascade;
drop table if exists collection_categories cascade;
drop table if exists collections cascade;
drop table if exists collection_collection_items cascade;
drop table if exists collection_workflow_transition_history cascade;
drop table if exists collection_metadata_items cascade;
drop
    type if exists collection_type cascade;
drop table if exists metadata_relationship cascade;
drop table if exists metadata_traits cascade;
drop table if exists metadata_categories cascade;
drop table if exists metadata_workflow_transition_history cascade;
drop table if exists metadata cascade;
drop
    type if exists metadata_type cascade;
drop table if exists traits cascade;
drop table if exists categories cascade;
drop
    type if exists metadata_status cascade;
drop
    type if exists metadata_status cascade;
drop table if exists workflow_states cascade;
drop
    type if exists workflow_state_type cascade;
-- +goose StatementEnd
