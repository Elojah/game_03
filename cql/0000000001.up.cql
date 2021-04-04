CREATE TABLE main.session (
    id              uuid PRIMARY KEY,
    user_id         uuid,
    twitch_token    varchar
);
