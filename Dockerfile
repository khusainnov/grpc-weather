FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o weather cmd/weather/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder app/weather .

EXPOSE 80

CMD ["/weather"]