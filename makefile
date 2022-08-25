generate:
	sqlc generate

engine:
	go build -o ${BINARY} .

serve:
	sudo docker-compose up --no-deps --build

build:
	docker build -t notes_app:latest .

run:
	make build
	sudo docker-compose -f dev-docker-compose.yaml up --no-deps --build

dev:
	sudo docker-compose -f dev-docker-compose.yaml up
