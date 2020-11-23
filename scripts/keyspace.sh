docker exec -it game_03_scylla cqlsh
CREATE KEYSPACE main WITH replication = {'class': 'SimpleStrategy','replication_factor' : 1} AND durable_writes = true;
