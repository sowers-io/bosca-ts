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
create table traits
(
    id   uuid    not null default gen_random_uuid(),
    name varchar not null,
    primary key (id)
);

create table categories
(
    id   uuid    not null default gen_random_uuid(),
    name varchar not null,
    primary key (id)
);

create type metadata_status as enum ('processing', 'ready');
create type metadata_type as enum ('standard', 'variant');

create table metadata
(
    id           uuid      not null default gen_random_uuid(),
    parent_id    uuid,
    name         varchar   not null,
    type         metadata_type      default 'standard',
    content_type varchar   not null,
    language_tag varchar   not null default 'en',
    tags         varchar[] not null,
    attributes   jsonb     not null,
    created      timestamp          default now(),
    modified     timestamp          default now(),
    status       metadata_status    default 'processing',
    primary key (id),
    foreign key (parent_id) references metadata (id) on delete cascade
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
drop table metadata_relationship cascade;
drop table metadata_traits cascade;
drop table metadata_categories cascade;
drop table metadata cascade;
drop table traits;
drop table categories;
-- +goose StatementEnd
