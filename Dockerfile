# build a golang image multistage

# build stage
FROM golang:alpine3.16 as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o notes_app .

# run stage
FROM alpine
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY --from=builder /app/notes_app .

EXPOSE 3000
CMD [ "/app/notes_app" ]