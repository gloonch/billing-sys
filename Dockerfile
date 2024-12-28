FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD ["./main"]
