FROM golang:latest

COPY ./ ./

RUN apt-get update

RUN go mod download
RUN go build -o main_service ./cmd/app/main.go

CMD ["./main_service"]