-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- articles
-- ------------------------------------------------------------
CREATE TABLE articles (
    id BIGINT PRIMARY KEY,
    title character varying(200) NOT NULL,
    content TEXT NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NULL
);
