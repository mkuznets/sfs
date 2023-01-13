CREATE TABLE episodes
(
    id          text primary key,
    channel_id  text references channels (id) not null,
    title       text                          not null,
    file_id     text references files (id)    not null,
    description text,
    link        text,
    authors     text,
    created_at  integer                       not null,
    updated_at  integer                       not null,
    deleted_at  integer
) STRICT;

CREATE INDEX episodes_channel_id_idx ON episodes (channel_id);

CREATE TABLE files
(
    id           text primary key,
    user_id      text references users (id) not null,
    url          text                       not null,
    size         integer                    not null CHECK (size > 0),
    content_type text                       not null,
    created_at   integer                    not null,
    updated_at   integer                    not null,
    deleted_at   integer
) STRICT;
