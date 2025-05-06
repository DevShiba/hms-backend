FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o hmsexe ./cmd/main.go

FROM scratch

COPY --from=builder /app/hmsexe /hmsexe
COPY --from=builder /app/.env /.env

EXPOSE 8080

ENTRYPOINT ["/hmsexe"]