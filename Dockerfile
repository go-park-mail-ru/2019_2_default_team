# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/
COPY logger/logger-config.json logger/logger-config.json
EXPOSE 8080
ENTRYPOINT ["/app/main"]