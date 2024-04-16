# Notes

`172.17.0.2` = IP Address for PostgreSQL from network settings from `docker container inspect postgres16`

`docker run --name simple-bank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.17.0.2:5432/simple_bank?sslmode=disable" simple-bank:latest` = start docker container and connect 2 stand-alone containers

Better way is to use user-defined network
`docker network ls`
`docker network create bank-network`
`docker network connect --help`
`docker network connect bank-network postgres16`
`docker container inspect postgres16` -> connected to two different networks -> bank-network and bridge(default)

`docker run --name simple-bank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres16:5432/simple_bank?sslmode=disable" simple-bank:latest`
