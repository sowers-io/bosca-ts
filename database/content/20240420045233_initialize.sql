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
create table models
(
    id            uuid    not null default gen_random_uuid(),
    type          varchar not null,
    name          varchar not null,
    description   varchar not null,
    configuration jsonb   not null,
    primary key (id)
);

insert into models (name, type, description, configuration)
values ('phi3:medium-128k', 'ollama-llm', 'phi3:medium-128k', '{"vectorSize": "5120", "contextWindow": "120000"}'::jsonb),
       ('phi3:mini', 'ollama-llm', 'phi3:mini', '{"vectorSize": "3072", "contextWindow": "8000"}'::jsonb);

create table prompts
(
    id          uuid    not null default gen_random_uuid(),
    name        varchar not null,
    description varchar not null,
    prompt      text    not null,
    primary key (id)
);

insert into prompts (name, description, prompt)
values ('Verse Labeler', '', $$You are a Biblical theologian who is also an expert in felt needs.  Your job is to classify verses, annotating them with felt needs.  Examples of the felt needs are:
* salvation
* forgiveness
* assurance
* deliverance
* healing
* provision
* protection
* guidance
* wisdom
* comfort
* peace
* joy
* love
* belonging
* purpose
* restoration
* vindication
* hope
* strength
* patience
* faith
* mercy
* humility
* holiness
* worship
You to add a column to the markdown table called `tags`, and fill that column with a set of felt needs that best apply to the given verse.

{table}
$$);

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
    configuration jsonb not null default '{"type": "default"}'::jsonb,
    primary key (system_id, model_id),
    foreign key (system_id) references storage_systems (id),
    foreign key (model_id) references models (id)
);

insert into storage_systems (name, description, type, configuration)
values ('Metadata Search Index', 'Index search content', 'search'::storage_system_type, '{"indexName": "metadata", "type": "meilisearch"}'::jsonb),
       ('Bible Chapter Vector Index', 'Store Bible Chapter Data', 'vector'::storage_system_type, '{"indexName": "biblechapters", "type": "qdrant", "vectorSize": "3072"}'::jsonb),
       ('Bible Verse Vector Index', 'Store Bible Verse Data', 'vector'::storage_system_type, '{"indexName": "bibleverses", "type": "qdrant", "vectorSize": "3072"}'::jsonb),
       ('Bible Verse Label Vector Index', 'Store Bible Verse Label Data', 'vector'::storage_system_type, '{"indexName": "bibleverselabels", "type": "qdrant", "vectorSize": "3072"}'::jsonb);

insert into storage_system_models (system_id, model_id)
values ((select id from storage_systems where name = 'Bible Chapter Vector Index'),
        (select id from models where name = 'phi3:mini'));

insert into storage_system_models (system_id, model_id)
values ((select id from storage_systems where name = 'Bible Verse Vector Index'),
        (select id from models where name = 'phi3:mini'));

create table workflows
(
    id            varchar not null, -- This is the identifier of the temporal workflow
    name          varchar not null,
    description   varchar not null,
    queue         varchar not null,
    configuration jsonb   not null default '{}',
    primary key (id)
);

create table workflow_activities
(
    id                   varchar not null,
    name                 varchar not null,
    description          varchar not null,
    child_workflow       boolean not null default false,
    child_workflow_queue varchar,
    configuration        jsonb   not null default '{}',
    primary key (id)
);

create type workflow_activity_parameter_type as enum ('context', 'supplementary', 'supplementary_array');

create table workflow_activity_inputs
(
    activity_id varchar                          not null,
    name        varchar                          not null,
    type        workflow_activity_parameter_type not null,
    primary key (activity_id, name),
    foreign key (activity_id) references workflow_activities (id) on delete cascade
);

create table workflow_activity_outputs
(
    activity_id varchar                          not null,
    name        varchar                          not null,
    type        workflow_activity_parameter_type not null,
    primary key (activity_id, name),
    foreign key (activity_id) references workflow_activities (id) on delete cascade
);

create table workflow_activity_instances
(
    id              bigserial not null,
    workflow_id     varchar   not null,
    activity_id     varchar   not null,
    execution_group int       not null,
    configuration   jsonb     not null default '{}',
    primary key (id)
);

create table workflow_activity_instance_inputs
(
    instance_id bigint  not null,
    name        varchar not null,
    value       jsonb   not null,
    primary key (instance_id, name),
    foreign key (instance_id) references workflow_activity_instances
);

create table workflow_activity_instance_outputs
(
    instance_id bigint  not null,
    name        varchar not null,
    value       jsonb   not null,
    primary key (instance_id, name),
    foreign key (instance_id) references workflow_activity_instances
);

create index ix_workflow_activity_instances_ix on workflow_activity_instances (workflow_id);

insert into workflows (id, name, description, queue, configuration)
values ('ProcessMetadata', 'Process Metadata', 'Process Metadata', 'metadata', '{"finalStateId": "draft"}'::jsonb),
       ('DeleteTemporary', 'Delete Temporary Metadata', 'Delete Temporary Metadata', 'metadata', '{}'::jsonb),
       ('ProcessTraits', 'Re-process traits', 'Process once they''ve already been processed', 'metadata', '{}'::jsonb),
       ('ProcessBible', 'Process Bible', 'Process Bible', 'bible', '{}'::jsonb),
       ('ProcessChapter', 'Process Chapter', 'Process Chapter', 'bible', '{}'::jsonb);

insert into workflow_activities (id, name, description)
values ('ProcessTraits', 'Process Metadata Traits', 'Execute workflows based on metadata traits'),
       ('ProcessSupplementaryTraits', 'Process Supplementary Metadata Traits', 'Execute workflows based on metadata supplementary traits'),
       ('GeneratePendingEmbeddingsFromTable', 'Generate Pending Embeddings from a Table', 'Generate Pending Embeddings from a Table, the supplied `column` will be the data'),
       ('GenerateChapters', 'Generate chapters', ''),
       ('GenerateChapterVerses', 'Generate verses', ''),
       ('GenerateChapterVerseTable', 'Generate Chapter Verse Table', 'Generate chapter verse table for the purposes of generating verse labels'),
       ('GenerateVerseLabelPendingEmbeddings', 'Generate Verse Label Pending Embeddings', ''),
       ('ExtractText', 'Extract Text from main content', 'Extract text from the main content'),
       ('ExecuteTablePrompt', 'Execute a table prompt', 'Execute a prompt by using the supplementary table data and save the results as supplementary data.  Uses context to leverage `inSupplementaryId` and `outSupplementaryId`'),
       ('GenerateTextEmbeddings', 'Generate Embeddings', 'Generate embeddings based on main content'),
       ('GenerateSupplementaryPendingEmbeddings', 'Generate Supplementary Embeddings', 'Generate embeddings based on supplementary content'),
       ('DeleteMetadata', 'Delete Metadata', 'Delete metadata'),
       ('DeleteSupplementary', 'Delete Supplementary Content', 'Delete supplementary Content, uses context with `supplementaryId` as the id to delete');

insert into workflow_activities (id, name, description, child_workflow, child_workflow_queue)
values ('AddToVectorIndex', 'Add to Vector Index', 'Add pending embeddings to vector index', true, 'vectors');
insert into workflow_activities (id, name, description, child_workflow, child_workflow_queue)
values ('AddToSearchIndex', 'Add to Search Index', 'Add text to search index', true, 'vectors');

insert into workflow_activity_inputs (activity_id, name, type)
values ('ExecuteTablePrompt', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('GeneratePendingEmbeddingsFromTable', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('GeneratePendingEmbeddingsFromTable', 'column', 'context'::workflow_activity_parameter_type),
       ('AddToVectorIndex', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('DeleteSupplementary', 'supplementaryIds', 'supplementary_array'::workflow_activity_parameter_type);

insert into workflow_activity_outputs (activity_id, name, type)
values ('GenerateChapterVerseTable', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('ExecuteTablePrompt', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('ExtractText', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type),
       ('GeneratePendingEmbeddingsFromTable', 'supplementaryId', 'supplementary'::workflow_activity_parameter_type);

insert into workflow_activity_instances (workflow_id, activity_id, execution_group)
values ('ProcessMetadata', 'ProcessTraits', 1),
       ('DeleteTemporary', 'DeleteMetadata', 1),
       ('ProcessBible', 'GenerateChapters', 1),
       ('ProcessChapter', 'GenerateChapterVerses', 1),
       ('ProcessChapter', 'GenerateChapterVerseTable', 1),
       ('ProcessChapter', 'ExecuteTablePrompt', 2),
       ('ProcessChapter', 'GeneratePendingEmbeddingsFromTable', 3),
       ('ProcessChapter', 'AddToVectorIndex', 4),
       ('ProcessChapter', 'DeleteSupplementary', 5);

insert into workflow_activity_instance_inputs (instance_id, name, value)
values ((select id from workflow_activity_instances where activity_id = 'ExecuteTablePrompt'), 'supplementaryId', '{"value": "chapter-verse-table"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'GeneratePendingEmbeddingsFromTable'), 'column', '{"value": "Labels"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'GeneratePendingEmbeddingsFromTable'), 'supplementaryId', '{"value": "chapter-verse-table-prompt-result"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'AddToVectorIndex'), 'supplementaryId', '{"value": "chapter-verse-pending-embeddings"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'DeleteSupplementary'), 'supplementaryIds', '{"values": ["chapter-verse-table", "chapter-verse-table-prompted", "chapter-verse-pending-embeddings"]}'::jsonb);

insert into workflow_activity_instance_outputs (instance_id, name, value)
values ((select id from workflow_activity_instances where activity_id = 'GenerateChapterVerseTable'), 'supplementaryId', '{"value": "chapter-verse-table"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'ExecuteTablePrompt'), 'supplementaryId', '{"value": "chapter-verse-table-prompt-result"}'::jsonb),
       ((select id from workflow_activity_instances where activity_id = 'GeneratePendingEmbeddingsFromTable'), 'supplementaryId', '{"value": "chapter-verse-pending-embeddings"}'::jsonb);

create table traits
(
    id          varchar not null,
    name        varchar not null,
    description varchar not null,
    primary key (id)
);

create table trait_workflows
(
    trait_id    varchar not null,
    workflow_id varchar not null,
    primary key (trait_id, workflow_id),
    foreign key (trait_id) references traits (id),
    foreign key (workflow_id) references workflows (id)
);

create table trait_workflow_activity_storage_systems
(
    trait_id          varchar not null,
    workflow_id       varchar not null,
    activity_id       varchar not null,
    storage_system_id uuid    not null,
    configuration     jsonb   not null default '{}'::jsonb,
    primary key (trait_id, workflow_id, activity_id, storage_system_id),
    foreign key (trait_id) references traits (id),
    foreign key (workflow_id) references workflows (id),
    foreign key (activity_id) references workflow_activities (id),
    foreign key (storage_system_id) references storage_systems (id)
);

create table trait_workflow_activity_prompts
(
    trait_id      varchar not null,
    workflow_id   varchar not null,
    activity_id   varchar not null,
    prompt_id     uuid    not null,
    configuration jsonb   not null default '{}'::jsonb,
    primary key (trait_id, workflow_id, prompt_id),
    foreign key (trait_id) references traits (id),
    foreign key (workflow_id) references workflows (id),
    foreign key (activity_id) references workflow_activities (id),
    foreign key (prompt_id) references prompts (id)
);

insert into traits (id, name, description)
values ('common.text', 'Textual Content', 'Generic Common Text'),
       ('bible.usx', 'Digital Bible', 'Generate vector and search indexed data.  Broken down by chapter and verse.'),
       ('bible.usx.chapter', 'Bible Chapter', 'USX Bible Chapter'),
       ('common.temporary', 'Temporary Content', 'Temporary Content');

insert into trait_workflows (trait_id, workflow_id)
values ('bible.usx', 'ProcessBible'),
       ('bible.usx.chapter', 'ProcessChapter'),
       ('common.temporary', 'DeleteTemporary');

insert into trait_workflow_activity_prompts (trait_id, workflow_id, activity_id, prompt_id, configuration)
values ('bible.usx.chapter', 'ProcessChapter', 'ExecuteTablePrompt', (select id from prompts where name = 'Verse Labeler'), '{}'::jsonb);

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
values ('processing', 'Processing', 'Initial Processing after Creation',
        'processing'::workflow_state_type, 'ProcessMetadata');

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
       ('draft', 'processing', 'reprocessing draft'),
       ('draft', 'published', 'publishing draft');

create type collection_type as enum ('root', 'standard', 'folder', 'queue');

create table collections
(
    id                        uuid      not null default gen_random_uuid(),
    name                      varchar   not null,
    type                      collection_type    default 'standard',
    attributes                jsonb     not null default '{}',
    labels                    varchar[] not null default '{}',
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
    trait_id      varchar,
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

create table sources
(
    id            uuid    not null default gen_random_uuid(),
    name          varchar not null,
    description   varchar not null,
    configuration jsonb   not null default '{}'::jsonb,
    primary key (id)
);

insert into sources (name, description)
values ('uploader', 'metadata from an upload using the uploader'),
       ('workflow', 'metadata generated during a workflow');

create type metadata_type as enum ('standard', 'variant');

create table metadata
(
    id                        uuid      not null                                  default gen_random_uuid(),
    parent_id                 uuid,
    name                      varchar   not null check (length(name) > 0),
    type                      metadata_type                                       default 'standard',
    content_type              varchar   not null check (length(content_type) > 0),
    content_length            bigint    not null check (content_length > 0),
    language_tag              varchar   not null check (length(language_tag) > 0) default 'en',
    labels                    varchar[] not null                                  default '{}',
    attributes                jsonb     not null                                  default '{}',
    created                   timestamp                                           default now(),
    modified                  timestamp                                           default now(),
    workflow_state_id         varchar   not null                                  default 'pending',
    workflow_state_pending_id varchar,
    metadata                  jsonb,
    source_id                 uuid,
    source_identifier         varchar,
    delete_workflow_id        varchar,
    primary key (id),
    foreign key (parent_id) references metadata (id) on delete cascade,
    foreign key (workflow_state_id) references workflow_states (id),
    foreign key (workflow_state_pending_id) references workflow_states (id),
    foreign key (source_id) references sources (id),
    foreign key (delete_workflow_id) references workflows (id)
);

create table metadata_supplementary
(
    metadata_id uuid      not null,
    key         varchar   not null,
    traits      varchar[] not null default '{}',
    created     timestamp not null default now(),
    modified    timestamp not null default now(),
    uploaded    timestamp,
    primary key (metadata_id, key)
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
    trait_id    varchar,
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
drop type if exists collection_type cascade;
drop table if exists metadata_relationship cascade;
drop table if exists metadata_traits cascade;
drop table if exists metadata_categories cascade;
drop table if exists metadata_workflow_transition_history cascade;
drop table if exists metadata cascade;
drop type if exists metadata_type cascade;
drop table if exists traits cascade;
drop table if exists categories cascade;
drop type if exists metadata_status cascade;
drop type if exists metadata_status cascade;
drop table if exists workflow_states cascade;
drop type if exists workflow_state_type cascade;
drop table if exists models cascade;
drop table if exists trait_workflows cascade;
drop table if exists workflow_trait_storage_systems cascade;
drop table if exists storage_systems cascade;
drop table if exists storage_system_models cascade;
drop type storage_system_type;
drop table prompts cascade;
drop table trait_workflow_activity_prompts cascade;
drop table trait_workflow_activity_storage_systems cascade;
drop table sources cascade;
drop table workflow_activities cascade;
drop table workflow_activity_instances cascade;
drop type workflow_activity_parameter_type cascade;
drop table metadata_supplementary cascade;
drop table workflow_activity_inputs cascade;
drop table workflow_activity_outputs cascade;
drop table workflow_activity_instance_inputs cascade;
drop table workflow_activity_instance_outputs cascade;
-- +goose StatementEnd
