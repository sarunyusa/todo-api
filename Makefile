local-run:
	@go run cmd/main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/todo-linux-x64 ./cmd/main.go

docker-db-up:
	docker run -d --rm --name todo-db -e POSTGRES_PASSWORD=P@ssw0rd -p 9990:5432 postgres

docker-db-down:
	docker stop todo-db