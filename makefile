generate:
	sqlc generate

run:
	go run ./app/main.go

engine:
	go build -o ${BINARY} app/*.go