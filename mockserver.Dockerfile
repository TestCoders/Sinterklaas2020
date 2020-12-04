# Build
FROM golang:alpine as builder
LABEL stage=builder

RUN apk add --no-cache gcc libc-dev
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/mockserver

# App
FROM alpine as application
WORKDIR /

COPY --from=builder /app/mockserver .

CMD [ "./mockserver" ]
