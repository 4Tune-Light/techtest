.PHONY: run migrate build docker-up docker-down

run:
	go run cmd/main.go

migrate:
	go run migrate/main.go

build:
	go build -o techtest cmd/main.go

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down -v

docker-clean:
	docker-compose down -v --rmi all
	docker system prune -f