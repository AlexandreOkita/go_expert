.PHONY: all migrate build_server build_client run_server run_client run_all clean

all: migrate build_server build_client run_all

migrate:
    docker-compose run --rm migration

build_server:
    go build -o server ./server/main.go

build_client:
    go build -o client ./client/main.go

run_server:
    ./server

run_client:
    ./client

run_all: run_server run_client

clean:
    rm -f server client