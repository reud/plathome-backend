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

Docker run 

- if WSL2 Docker

```shell script
docker run --name be -e host=10.255.0.1 -d --rm --network host reud/plathome-backend
```

