CREATE TABLE humans
(
    id           serial
        constraint humans_pk primary key,
    name         VARCHAR(256) DEFAULT 'Anonymous' NOT NULL,
    sex          SMALLINT NOT NULL,
    favorability SMALLINT     DEFAULT 0 NOT NULL,
    description  TEXT,
    created_at   DATE     NOT NULL,
    update_at    DATE,
    deleted_at   DATE
);

