FROM golang:latest

WORKDIR /app
COPY go.mod ./
RUN go mod tidy

COPY . .
RUN go build -o stress ./cmd/stress

ENTRYPOINT ["./stress"]