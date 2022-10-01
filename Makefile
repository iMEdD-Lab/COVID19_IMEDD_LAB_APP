# migrations for local db usage
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host --user $(id -u):$(id -g) migrate/migrate -path=/migrations/ -database 'postgres://127.0.0.1:5433/covid19?sslmode=disable&user=admin&password=password'
MIGRATE_CREATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host --user $(shell id -u):$(shell id -g) migrate/migrate create --seq -ext sql -dir /migrations/

.PHONY: populate-db
populate-db: ## populate the database
	go run cmd/populate-db/main.go

.PHONY: db-start
db-start: ## start the database
	@mkdir -p testdata/postgres
	docker run --rm --name covid19db -d -v $(shell pwd)/testdata:/testdata -p 5433:5432 \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=password -e POSTGRES_DB=covid19 -e POSTGRES_USER=admin -d postgres:12.3

.PHONY: db-stop
db-stop: ## stop the database
	docker stop covid19db

.PHONY: db-login
db-login: ## login to the database
	docker exec -it covid19db psql -U admin -d covid19

.PHONY: migrate
migrate: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE_CREATE) $${name}
