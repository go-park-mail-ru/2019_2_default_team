# build stage
FROM golang as builder

ENV GO111MODULE=on

RUN ls /go/src
RUN cd /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o session session_main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/session /app/
#COPY ../logger/logger-config.json logger/logger-config.json
EXPOSE 8081
ENTRYPOINT ["/app/session"]