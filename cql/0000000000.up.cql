-- USER MANAGEMENT
-- ---------------

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
    token           varchar
);

CREATE INDEX ON main.session(user_id);

-- ROOM MANAGEMENT
-- ---------------

CREATE TABLE main.room (
    id              uuid PRIMARY KEY,
    owner_id        uuid,
    world_id        uuid,
    name            varchar
);

CREATE INDEX ON main.room(owner_id);

CREATE TABLE main.room_user (
    user_id        uuid,
    room_id        uuid,
    role           tinyint,
    PRIMARY KEY(user_id, room_id)
);

CREATE INDEX ON main.room_user(room_id);

-- ENTITY MANAGEMENT
-- -----------------

-- must replicate main.entity_backup fields
CREATE TABLE main.entity (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    cell_id         uuid,
    name            varchar,
    x               bigint,
    y               bigint,
    rot             smallint,
    radius          smallint,
    at              bigint,
    animation_id    uuid,
    animation_at    bigint
);

CREATE INDEX ON main.entity(cell_id);

-- must replicate main.entity fields
CREATE TABLE main.entity_backup (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    cell_id         uuid,
    name            varchar,
    x               bigint,
    y               bigint,
    rot             smallint,
    radius          smallint,
    at              bigint,
    animation_id    uuid,
    animation_at    bigint
);

CREATE TABLE main.entity_animation (
    id              uuid PRIMARY KEY,
    entity_id       uuid,
    sheet_id        uuid,
	duplicate_id    uuid, -- id for client loading
    name            varchar,
    start           bigint,
    end             bigint,
    sequence        list<bigint>,
    rate            bigint,
    frame_width     bigint,
    frame_height    bigint,
    frame_start     bigint,
    frame_end       bigint,
    frame_margin    bigint,
    frame_spacing   bigint
);

CREATE INDEX ON main.entity_animation(entity_id);

CREATE TABLE main.entity_template (
    id              uuid PRIMARY KEY,
    name            varchar
);

CREATE INDEX ON main.entity_template(name);

-- CHARACTER

CREATE TABLE main.pc (
    id              uuid,
    user_id         uuid,
    world_id        uuid,
    entity_id       uuid,
    PRIMARY KEY(user_id, world_id, id)
);

CREATE TABLE main.npc (
    id              uuid,
    world_id        uuid,
    entity_id       uuid,
    PRIMARY KEY(world_id, id)
);

-- WORLD MANAGEMENT
-- ----------------

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
    tilemap         uuid,
    x               bigint,
    y               bigint
);
