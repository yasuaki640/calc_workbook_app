CREATE TABLE go_crud.humans
(
    id           serial
        constraint humans_pk primary key,
    name         VARCHAR(256) DEFAULT 'Anonymous' NOT NULL,
    sex          SMALLINT  NOT NULL,
    birthday     DATE,
    description  TEXT,
    favorability SMALLINT     DEFAULT 0 NOT NULL,
    created_at   TIMESTAMP NOT NULL,
    update_at    TIMESTAMP,
    deleted_at   TIMESTAMP
);

alter table go_crud.humans
    owner to go_crud;