BEGIN;

CREATE TABLE go_crud.posts
(
    id       SERIAL
        CONSTRAINT posts_pk
            PRIMARY KEY,
    human_id INT
        CONSTRAINT posts_humans_id_fk
            REFERENCES humans
            ON UPDATE CASCADE ON DELETE CASCADE,
    caption  TEXT
);
ALTER TABLE go_crud.posts
    OWNER TO go_crud;

COMMIT;