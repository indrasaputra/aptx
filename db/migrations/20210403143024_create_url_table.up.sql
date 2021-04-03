BEGIN;

CREATE TABLE IF NOT EXISTS urls (
    id              BIGSERIAL   PRIMARY KEY,
    code            VARCHAR(20) NOT NULL UNIQUE,
    short_url       TEXT        NOT NULL UNIQUE,
    original_url    TEXT        NOT NULL,
    expired_at      TIMESTAMP,
    created_at      TIMESTAMP
);

COMMIT;