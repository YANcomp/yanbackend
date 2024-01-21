-- +goose Up
CREATE
EXTENSION IF NOT EXISTS citext;
CREATE
EXTENSION IF NOT EXISTS POSTGIS;
CREATE
EXTENSION IF NOT EXISTS pg_trgm;
CREATE
EXTENSION IF NOT EXISTS btree_gist;
DROP TABLE IF EXISTS stories CASCADE;
DROP TABLE IF EXISTS slides CASCADE;

create table stories
(
    id             serial PRIMARY KEY,
    uuid           UUID          DEFAULT gen_random_uuid(),
    preview        text                     DEFAULT '',
    title          VARCHAR(250) NOT NULL CHECK ( title <> '' ),
    isActive       boolean                  DEFAULT false,
    isActiveMobile boolean                  DEFAULT false,
    created_at     TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at     TIMESTAMP WITH TIME ZONE
);

create table slides
(
    id                 serial PRIMARY KEY,
    uuid               UUID                     DEFAULT gen_random_uuid(),
    delay              integer                  DEFAULT 0,
    caption            text                     DEFAULT '',
    content            text                     DEFAULT '',
    textPosition       text                     DEFAULT '',
    backgroundImage    text                     DEFAULT '',
    isHideShadowBottom boolean                  DEFAULT false,
    story_id           integer NOT NULL,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at         TIMESTAMP WITH TIME ZONE
);

-- +goose Down
drop table stories;
drop table slides;