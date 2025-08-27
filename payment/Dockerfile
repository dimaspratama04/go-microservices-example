FROM golang:1.24-alpine AS builder

WORKDIR /api

ADD go.mod go.sum ./

COPY . .

RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /api/payment-service main.go

FROM alpine AS runtime

WORKDIR /api

COPY --from=builder /api/payment-service .

COPY --from=builder /api/ .

ENTRYPOINT [ "/api/payment-service" ]
