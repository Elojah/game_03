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

CREATE TABLE main.session (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    twitch_token    varchar
);

CREATE INDEX ON main.session(user_id);

CREATE TABLE main.room (
    id              uuid PRIMARY KEY,
    owner_id        uuid,
    world_id        uuid
);

CREATE INDEX ON main.room(owner_id);

CREATE TABLE main.pc (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    room_id         uuid
);

CREATE INDEX ON main.pc(user_id);
CREATE INDEX ON main.pc(room_id);

CREATE TABLE main.entity (
    id              uuid PRIMARY KEY,
    pc_id           uuid,
    x               bigint,
    y               bigint,
    rot             smallint,
    radius          smallint
);

CREATE TABLE main.world (
    id              uuid PRIMARY KEY,
    height          bigint,
    width           bigint,
    field           map<smallint, uuid>
);

CREATE TABLE main.cell (
    world_id        uuid,
    x               bigint,
    y               bigint,
    field           blob,
    PRIMARY KEY(world_id, x, y)
);
