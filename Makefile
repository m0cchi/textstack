

run-sqlc:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate -f sqlc.yml

migrate-local-postgres:
	migrate -database "postgresql://localhost:5432/textstack?user=textstack&password=textstack&sslmode=disable" -path misc/db/migrations up



