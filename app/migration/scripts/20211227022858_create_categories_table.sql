-- +goose Up
-- +goose StatementBegin
create table if not exists categories (
    id          varchar(6)   not null primary key,
    parent_id   varchar(6),
    name        varchar(255) not null,
    description text         not null,
    slug        varchar(64)  not null,
    page_title  varchar(255) not null,
    meta        jsonb        not null,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),
    deleted_at  timestamptz,
    foreign key (parent_id) references categories (id) on delete cascade
);
create unique index if not exists ux_categories_slug on categories (slug);
create index if not exists ix_categories_parent_id  on categories (parent_id);
create index if not exists ix_categories_created_at on categories (created_at);
create index if not exists ix_categories_updated_at on categories (updated_at);
create index if not exists ix_categories_deleted_at on categories (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists ix_categories_deleted_at;
drop index if exists ix_categories_updated_at;
drop index if exists ix_categories_created_at;
drop index if exists ix_categories_parent_id;
drop index if exists ux_categories_slug;
drop table if exists categories;
-- +goose StatementEnd
