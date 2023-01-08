# AgoraCore
The core lib that empowers others to build Agoras (ancient marketplaces).


# POSTGRES
```sh
docker run --name postgres --env=POSTGRES_PASSWORD=password --env=PGDATA=/var/lib/postgresql/data --volume=/var/lib/postgresql/data -p 5432:5432 -d postgres:15
```


# Versions
v0.0.0 - Empty 

v0.0.1 - Initial data modeling

v0.0.2 - Initial db actions

---
docker run --name=teampostgresql --env=TEAMPOSTGRESQL_ADMIN_USER=postgres --env=TEAMPOSTGRESQL_ADMIN_PASSWORD=password -p 8082:8082 -d teampostgresql/teampostgresql
