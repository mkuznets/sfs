CREATE TABLE channels
(
    id          text primary key,
    user_id     text references users (id),
    title       text,
    link        text,
    authors     text,
    description text,
    created_at  integer not null,
    updated_at  integer,
    deleted_at  integer
) STRICT;

CREATE INDEX channels_user_id_idx ON channels (user_id);
CREATE UNIQUE INDEX channels_title_unique ON channels (user_id, title) WHERE deleted_at IS NULL;

CREATE TABLE users
(
    id    text primary key,
    email text
) STRICT;
