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
    world_id        uuid,
    name            varchar
);

CREATE INDEX ON main.room(owner_id);

CREATE TABLE main.pc (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    world_id        uuid,
    entity_id       uuid
);

CREATE INDEX ON main.pc(user_id);
CREATE INDEX ON main.pc(world_id);

CREATE TABLE main.entity (
    id              uuid PRIMARY KEY,
    cell_id         uuid,
    x               bigint,
    y               bigint,
    rot             smallint,
    radius          smallint,
    at              bigint,
);

CREATE INDEX ON main.entity(cell_id);

CREATE TABLE main.world (
    id              uuid PRIMARY KEY,
    width           bigint,
    height          bigint,
    cell_width      bigint,
    cell_height     bigint
);

CREATE TABLE main.world_cell (
    world_id        uuid,
    cell_id         uuid,
    x               bigint,
    y               bigint,
    PRIMARY KEY(world_id, x, y)
);

CREATE TABLE main.cell (
    id              uuid PRIMARY KEY,
    contiguous      map<tinyint, uuid>,
    tilemap         map<bigint, uuid>
);
