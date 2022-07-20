FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build ./cmd/feedback_service

EXPOSE 9000

CMD ["./feedback_service"]

