local-run:
	@go run cmd/main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/todo-linux-x64 ./cmd/main.go

api-docker-image:  build
	docker build -t todo-api -f ./scripts/todo.Dockerfile .

start-compose:
	docker-compose -f ./scripts/todo-compose.yml up -d

stop-compose:
	docker-compose -f ./scripts/todo-compose.yml down

restart-compose:
	docker-compose -f ./scripts/todo-compose.yml restart

log-compose:
	docker-compose -f ./scripts/todo-compose.yml logs -f