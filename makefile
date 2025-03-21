DB_DSN := "postgres://postgres:1234567@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run D:\RestAPI\P-ex.7\cmd\app\main.go

gen tsk:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
gen usr:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number