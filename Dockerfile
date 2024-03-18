### build ###

FROM golang:alpine AS builder
WORKDIR /app

# dependencies
COPY go.mod go.sum ./
RUN go mod download
# app
COPY . .

RUN go build -o zumm .

### deploy ### 

FROM debian:bookworm
WORKDIR /root/
COPY --from=builder /app/zumm .
EXPOSE 8080

### run ###

CMD ["./zumm"]
