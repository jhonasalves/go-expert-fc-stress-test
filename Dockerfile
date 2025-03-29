FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o loadtester .

FROM scratch

COPY --from=builder /app/loadtester .

ENTRYPOINT ["./loadtester"]

