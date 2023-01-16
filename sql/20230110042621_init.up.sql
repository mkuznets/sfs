create table channels
(
    id                text primary key           not null check (substring(id, 1, 3) = 'ch_'),
    user_id           text references users (id) not null,
    title             text                       not null,
    link              text                       not null,
    authors           text                       not null,
    description       text                       not null,

    feed_content      blob,
    feed_url          text                       not null,
    feed_published_at integer check (feed_published_at is null or feed_published_at > 0),

    created_at        integer                    not null check (created_at > 0),
    updated_at        integer                    not null check (updated_at > 0),
    deleted_at        integer check (deleted_at is null or deleted_at > 0)
) strict;

create index channels_user_id_idx on channels (user_id);
create unique index channels_title_unique on channels (user_id, title) where deleted_at is null;

create table users
(
    id    text primary key not null check (substring(id, 1, 4) = 'usr_'),
    email text
) strict;

create table episodes
(
    id          text primary key              not null check (substring(id, 1, 3) = 'ep_'),
    channel_id  text references channels (id) not null,
    title       text                          not null,
    file_id     text references files (id)    not null,
    description text                          not null,
    link        text                          not null,
    authors     text                          not null,
    created_at  integer                       not null check (created_at > 0),
    updated_at  integer                       not null check (updated_at > 0),
    deleted_at  integer
) strict;

create index episodes_channel_id_idx on episodes (channel_id);

create table files
(
    id                 text primary key           not null check (substring(id, 1, 5) = 'file_'),
    user_id            text references users (id) not null,
    url                text                       not null,
    size               integer                    not null check (size > 0),
    content_type text                       not null,
    created_at         integer                    not null check (created_at > 0),
    updated_at         integer                    not null check (updated_at > 0),
    deleted_at         integer
) strict;
