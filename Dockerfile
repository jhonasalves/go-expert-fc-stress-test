FROM golang:1.23-alpine as builder

WORKDIR /app

RUN apk update && apk add ca-certificates

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o loadtester .

FROM scratch

COPY --from=builder /app/loadtester .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./loadtester"]