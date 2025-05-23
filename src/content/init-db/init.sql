CREATE TABLE IF NOT EXISTS entries
(
    id             INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title          TEXT        NOT NULL,
    description    TEXT        NOT NULL,
    author         VARCHAR(32) NOT NULL,
    created_at     TIMESTAMP   NOT NULL,
    last_edited_at TIMESTAMP   NOT NULL,
    is_private     boolean     NOT NULL default true,
    tags           TEXT[]      NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS comments
(
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    post_id    INT,
    author     VARCHAR(32) NOT NULL,
    text       TEXT        NOT NULL,
    created_at TIMESTAMP   NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
