run:
	go run main.go

up:
	docker compose up -d

down:
	docker compose down

remove:
	rm -r volumes

all:
	make down
	make up
	make run

restart:
	make down
	make remove
	make up
	make run

test:
	go test ./... -v -cover

swag-init:
	${GOPATH}/bin/swag init --md ./

swag-format:
	${GOPATH}/bin/swag fmt

swag:
	make swag-format
	make swag-init

format:
	go fmt ./...