DROP TABLE IF EXISTS humans;

CREATE TABLE humans
(
    id         integer primary key,
    name       text,
    sex        integer,
    status     text,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);

INSERT INTO humans
VALUES (1, 'Humika Baba', 2, 'In a relationship', datetime('now'), datetime('now'), NULL);
INSERT INTO humans
VALUES (2, 'Mio Imada', 2, 'Been with married', datetime('now'), datetime('now'), NULL);
INSERT INTO humans
VALUES (3, 'Rika Izumi', 2, 'Lover', datetime('now'), datetime('now'), NULL);
INSERT INTO humans
VALUES (4, 'Kayako Abe', 2, 'Girl friend', datetime('now'), datetime('now'), NULL);
