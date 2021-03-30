CREATE TABLE main.user (
    id              uuid PRIMARY KEY,
    twitch_id       varchar
);

CREATE INDEX ON main.user(twitch_id);

CREATE TABLE main.role (
    id              uuid PRIMARY KEY,
    set_id          uuid,
    name            varchar
);

CREATE INDEX ON main.role(set_id);
