FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o microservice .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/microservice .
EXPOSE 8080
CMD ["./microservice"]
