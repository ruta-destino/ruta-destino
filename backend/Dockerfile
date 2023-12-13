FROM golang:1.21-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o backend ./main.go

FROM alpine:3.17
COPY --from=builder /app/backend /usr/local/bin/backend
CMD [ "backend" ]
