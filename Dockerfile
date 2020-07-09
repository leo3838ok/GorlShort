FROM golang:1.14.4-alpine3.12 AS build
WORKDIR /src
COPY . .
RUN go build -mod=vendor -o bin/server .

FROM alpine:3.12
COPY --from=build /src/bin/server .
ENTRYPOINT ./server