generate:
	sqlc generate

engine:
	go build -o ${BINARY} app/*.go

serve:
	sudo docker-compose up --no-deps --build

build:
	docker build -t notes_app:latest .

run:
	make build
	sudo docker-compose -f dev-docker-compose.yaml up --no-deps --build

