FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go mod download
RUN go build ./cmd/feedback_service

EXPOSE 9000

CMD goose --dir migrations postgres "user=postgres password=postgres host=postgres dbname=test sslmode=disable" up; ./feedback_service

