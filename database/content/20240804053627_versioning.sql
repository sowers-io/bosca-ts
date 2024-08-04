-- +goose Up
-- +goose StatementBegin
alter table metadata
    add column version int not null default 1;
alter table metadata
    add column active_version int not null default 1;

create table metadata_versions
(
    id                        uuid      not null,
    version                   int       not null,
    parent_id                 uuid,
    name                      varchar   not null check (length(name) > 0),
    type                      metadata_type                                       default 'standard',
    content_type              varchar   not null check (length(content_type) > 0),
    content_length            bigint,
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
    primary key (id, version),
    foreign key (id) references metadata (id) on delete cascade,
    foreign key (parent_id) references metadata (id) on delete cascade,
    foreign key (source_id) references sources (id)
);

create table metadata_version_traits
(
    metadata_id uuid,
    version     int,
    trait_id    varchar,
    primary key (metadata_id, version, trait_id),
    foreign key (metadata_id, version) references metadata_versions (id, version) on delete cascade,
    foreign key (trait_id) references traits (id) on delete cascade
);

create table metadata_version_categories
(
    metadata_id uuid,
    version     int,
    category_id uuid,
    primary key (metadata_id, version, category_id),
    foreign key (metadata_id, version) references metadata_versions (id, version) on delete cascade,
    foreign key (category_id) references categories (id) on delete cascade
);

create or replace function metadata_versions_trigger() returns trigger AS
$version_trigger$
begin
    if new.version = old.version then
        return new;
    end if;

    insert into metadata_versions (id, version, parent_id, name, content_type, content_length,
                                   workflow_state_pending_id, metadata, source_id, source_identifier,
                                   delete_workflow_id)
    values (old.id, old.version, old.parent_id, old.name, old.content_type, old.content_length,
            old.workflow_state_pending_id, old.metadata, old.source_id, old.source_identifier,
            old.delete_workflow_id);

    insert into metadata_version_traits (metadata_id, version, trait_id)
    select metadata_id, new.version, trait_id
    from metadata_traits
    where metadata_id = new.id;

    delete from metadata_traits where metadata_id = new.id;

    insert into metadata_version_categories (metadata_id, version, category_id)
    select metadata_id, new.version, category_id
    from metadata_categories
    where metadata_id = new.id;

    delete from metadata_categories where metadata_id = new.id;

    return new;
end;
$version_trigger$ language plpgsql;

create trigger metadata_versions
    before update
    on metadata
    for each row
execute function metadata_versions_trigger();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists metadata_versions cascade;
drop table if exists metadata_version_traits cascade;
drop table if exists metadata_version_categories cascade;
drop function metadata_versions_trigger() cascade;
-- +goose StatementEnd
