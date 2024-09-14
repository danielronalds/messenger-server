build:
	docker compose up --build api

build-all:
	docker compose up --build

dev:
	go run .

run-db:
	docker compose start db

stop:
	docker compose stop

clean:
	docker compose down -v

fmt:
	go fmt github.com/danielronalds/...

test:
	go test github.com/danielronalds/...

ls:
	cat Makefile
