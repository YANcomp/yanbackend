-- +goose Up
BEGIN;

SET
    statement_timeout = 0;
SET
    client_encoding = 'UTF8';
SET
    standard_conforming_strings = ON;
SET
    check_function_bodies = FALSE;
SET
    client_min_messages = WARNING;
SET
    search_path = public, extensions;
SET
    default_tablespace = '';
SET
    default_with_oids = FALSE;

-- EXTENSIONS --
CREATE
    EXTENSION IF NOT EXISTS pgcrypto;
CREATE
    EXTENSION IF NOT EXISTS citext;

-- DROP TABLE --
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

COMMIT;

-- +goose Down
drop table stories;
drop table slides;