

run-sqlc:
	docker run --rm -u=$$UID:$$UID  -v $$(pwd):/src -w /src kjconroy/sqlc generate -f sqlc.yml

build-sqlc-grpc:
	docker build misc/rpc/sqlc_grpc -t sqlc_grpc

run-sqlc-grpc: build-sqlc-grpc
	docker run --rm -u=$$UID:$$UID -v $$(pwd):/src -w /src sqlc_grpc

migrate-local-postgres:
	migrate -database "postgresql://localhost:5432/textstack?user=textstack&password=textstack&sslmode=disable" -path misc/db/migrations up



