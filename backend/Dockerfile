FROM golang:1.26.0 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend ./cmd/api

FROM alpine:3.18.12

WORKDIR /app

COPY --from=build /app/config.yaml .
COPY --from=build /app/secret ./secret/

COPY --from=build /app/backend ./backend

RUN addgroup -S appgroup && adduser -S appuser -G appgroup && chmod +x ./backend && chown -R appuser:appgroup /app

USER appuser

ENTRYPOINT ["./backend"]
