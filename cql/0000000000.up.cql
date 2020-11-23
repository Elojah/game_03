CREATE TABLE main.user (
    id              uuid PRIMARY KEY,
    email           varchar,
    password        varchar,
    firebase_id     varchar
);

CREATE INDEX ON main.user(email);
