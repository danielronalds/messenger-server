# Lists all Make commands
ls:
	@grep "^	" -v Makefile | grep "shell" -v | grep "ifdef" -v | grep "else" -v | grep "endif" -v

# Build the API docker container
build:
	docker compose up --build api

# Build all the docker container
build-all:
	docker compose up --build

# Run go locally
dev:
	go run .

# Start only the postgres docker container
run-db:
	docker compose start db

# Stop docker containers
stop:
	docker compose stop

# Clean the docker containers, useful for resetting the db
clean:
	docker compose down -v

# Formats the code
fmt:
	go fmt github.com/danielronalds/...

GOTESTSUM_VERSION := $(shell gotestsum --version 2>/dev/null)

# Runs all tests
test:
ifdef GOTESTSUM_VERSION
	gotestsum --format testname github.com/danielronalds/messenger-server/...
else
	go test github.com/danielronalds/messenger-server/...
endif

# Seeds the database with some default data
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
