FROM golang:1.26.0 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend ./cmd/api

FROM alpine:3.18.12

USER nobody

WORKDIR /app

COPY --from=build /app/config.yaml .
COPY --from=build /app/secret .

COPY --from=build /app/backend ./backend

ENTRYPOINT ["./backend"]
