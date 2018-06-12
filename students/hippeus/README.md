# Notes

This exercise has an external dependency to a database. The solution handles it by encapsulating database within docker container. Such design has its trade-off as in order to make it runnable you need docker and docker-compose installed on your host machine.

## Usage

### Run an instance of a container using docker-compose

    `docker-compose -f psql-init.yml up -d`

### Verify that database stores correct input records

    `docker exec -it <container_name> psql -d phone -U admin`
    (inside container) `select * from phone_numbers;`

### Install and run phone app

    `go install ./cmd/phone.go`

### Verify results by inspecting postgres database

    `docker exec -it <container_name> psql -d phone -U admin`
    (inside container) `select * from phone_numbers;`

### Kill container

    `docker-compose -f psql-init.yml down`

### Rinse and repeat

    NOTE: In order to repeat this exercise, you might need to restore original entries in a database. To do so, simply kill your container instance and run this command:
    `docker volume rm hippeus_psql-data`