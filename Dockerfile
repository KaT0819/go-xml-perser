# Build Stage
FROM golang:1.16 As builder
WORKDIR /app

COPY . .

ENV GO111MODULE on
COPY go.mod go.sum ./
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o /main ./src/main.go

# Runtime Stage
FROM alpine:3.10.2
RUN apk add --no-cache ca-certificates

ENV SPANNER_PROJECT_ID <GCPプロジェクトID>
ENV SPANNER_INSTANCE_ID test-instance
ENV SPANNER_DATABASE_ID test-database

COPY --from=builder /main .
CMD ["./main", "--host", "0.0.0.0", "--port", "8080"]

EXPOSE 8080