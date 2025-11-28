
FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go build -o orders ./cmd/server

FROM debian:bookworm
COPY --from=build /app/orders /usr/local/bin/orders
CMD ["orders"]
