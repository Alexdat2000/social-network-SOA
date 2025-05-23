CREATE TABLE IF NOT EXISTS users
(
    username        VARCHAR(100) UNIQUE NOT NULL PRIMARY KEY,
    email           VARCHAR(100)        NOT NULL,
    hashed_password VARCHAR(64)         NOT NULL,
    first_name      VARCHAR(100),
    last_name       VARCHAR(100),
    date_of_birth   DATE,
    phone_number    varchar(15),
    created_at      TIMESTAMP           NOT NULL,
    last_edited_at  TIMESTAMP           NOT NULL
);
