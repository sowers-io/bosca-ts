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
create table profile_attribute_types
(
    id          uuid    not null,
    name        varchar not null,
    description varchar not null,
    primary key (id)
);

create type profile_visibility as enum ('system', 'user', 'friends', 'friends_of_friends', 'public');

create table profiles
(
    id         uuid               not null default gen_random_uuid(),
    principal  varchar            not null check (length(principal) > 0),
    name       varchar            not null check (length(name) > 0),
    visibility profile_visibility not null default 'system'::profile_visibility,
    created    timestamp          not null default now(),
    primary key (id)
);

comment on column profiles.principal is 'this is the identity provider id';

create table profile_attributes
(
    id         uuid               not null,
    profile_id uuid               not null,
    type_id    uuid               not null,
    visibility profile_visibility not null default 'system'::profile_visibility,
    value_type varchar            not null,
    value      bytea              not null,
    confidence float              not null,
    priority   int                not null,
    source     varchar            not null,
    created    timestamp          not null default now(),
    expiration timestamp,
    primary key (id),
    foreign key (profile_id) references profiles (id) on delete cascade,
    foreign key (type_id) references profile_attribute_types (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table profile_attributes cascade;
drop table profiles cascade;
drop type profile_visibility cascade;
drop table profile_attribute_types cascade;
-- +goose StatementEnd
