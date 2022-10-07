-- NPC MANAGEMENT
-- -----------------

CREATE TABLE main.npc (
    id              uuid,
    world_id        uuid,
    entity_id       uuid,
    PRIMARY KEY(world_id, id)
);
