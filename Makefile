

run-sqlc:
	docker run --rm -u=$$UID:$$UID  -v $$(pwd):/src -w /src kjconroy/sqlc generate -f sqlc.yml
	python misc/rpc/generator/bin/generator.py tmp/model.json proto/textstack.proto Textstack textstack_rpc


migrate-local-postgres:
	migrate -database "postgresql://localhost:5432/textstack?user=textstack&password=textstack&sslmode=disable" -path misc/db/migrations up



