Backend
====

Postgres Setting

```shell script
docker run --name plathome-db -p 5432:5432 -d postgres:11.5
```

Access To Postgres
```shell script
docker exec -it plathome-db psql -U postgres
```