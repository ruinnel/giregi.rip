FROM golang:1.15 AS builder

COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/server app/main.go

FROM gruebel/upx:latest as upx
COPY --from=builder /src/build/server /server.origin
RUN upx --best --lzma -o /server /server.origin

FROM alpine:latest AS product-stage
RUN mkdir /app
COPY --from=upx /server /app/server
COPY ./.sql-migrate /app/.sql-migrate
WORKDIR /app
