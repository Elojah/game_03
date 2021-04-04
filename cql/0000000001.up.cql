CREATE TABLE main.session (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    twitch_token    varchar
);

CREATE INDEX ON main.session(user_id);

CREATE TABLE main.room (
    id              uuid PRIMARY KEY,
    owner_id        uuid
);

CREATE INDEX ON main.room(owner_id);

CREATE TABLE main.pc (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    room_id         uuid
);

CREATE INDEX ON main.pc(user_id);
CREATE INDEX ON main.pc(room_id);
