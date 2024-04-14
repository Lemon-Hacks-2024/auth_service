FROM golang:latest

COPY ./ ./

RUN apt-get update

RUN go mod download
RUN go build -o task_tracking_service ./cmd/app/main.go

CMD ["./main_service"]