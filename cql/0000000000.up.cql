CREATE TABLE main.user (
    id              uuid PRIMARY KEY,
    email           varchar,
    password        varchar,
    firebase_id     varchar
);

CREATE INDEX ON main.user(email);

CREATE TABLE main.role (
    id              uuid PRIMARY KEY,
    set_id          uuid,
    name            varchar
);

CREATE INDEX ON main.role(set_id);
