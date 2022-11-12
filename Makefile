
run:
	- go mod tidy
	- go run cmd/main.go

db:
	- docker-compose -f docker-compose.yaml up -d

migration:
	- migrate -path ./schema -database 'postgres://postgres:bloggy@localhost:5437/postgres?sslmode=disable' up

migration-down:
	- migrate -path ./schema -database 'postgres://postgres:bloggy@localhost:5437/postgres?sslmode=disable' down

db-down:
	- docker-compose -f docker-compose.yaml down

inside-db:
	- docker exec -it bloggy-db /bin/bash