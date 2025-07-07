# Dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

CMD ["/app/main"]
