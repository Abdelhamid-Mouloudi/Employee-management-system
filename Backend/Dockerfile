FROM golang:1.23.4 AS builder

WORKDIR /backend

COPY . .

RUN go mod tidy


RUN go build -o /backend/main main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /backend

COPY --from=builder /backend/main .

EXPOSE 8080

CMD ["./main"]

