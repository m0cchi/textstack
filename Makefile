
setup-tools:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

	@echo "Show https://github.com/golang-migrate/migrate/tree/master/cmd/migrate"

run-sqlc:
	docker run --rm -u=$$UID:$$UID  -v $$(pwd):/src -w /src kjconroy/sqlc generate -f sqlc.yml

tmp/model.json:
	make run-sqlc

generate-protobuf: tmp/model.json
	python misc/rpc/generator/bin/generator.py tmp/model.json proto/textstack.proto Textstack
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $$(find proto -name "*proto" -type f)


migrate-local-postgres:
	migrate -database "postgresql://localhost:5432/textstack?user=textstack&password=textstack&sslmode=disable" -path misc/db/migrations up



