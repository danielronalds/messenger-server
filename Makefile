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

seed-db:
	docker compose up -d
	@echo "Waiting for containers to spin up!"
	@sleep 10
	@echo ""
	curl -X POST -H "Content-Type: application/json" -d '{"username":"johns","displayname":"John Smith","password":"pass"}' http://localhost:8080/users
	@echo ""
	@echo ""
	curl -X POST -H "Content-Type: application/json" -d '{"username":"janes","displayname":"Jane Smith","password":"password"}' http://localhost:8080/users
	@echo ""
	@echo ""
	curl -X POST -H "Content-Type: application/json" -d '{"username":"jonsnow","displayname":"Jon Snow","password":"winterIsComing"}' http://localhost:8080/users


ls:
	cat Makefile
