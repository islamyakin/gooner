FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY *.mod .
RUN go mod download
RUN go mod verify

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /gooner
FROM gcr.io/distroless/base-debian11
COPY --from=builder /gooner /gooner
ENTRYPOINT ["/gooner"]