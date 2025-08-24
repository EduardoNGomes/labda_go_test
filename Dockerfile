FROM golang:1.25 AS build
WORKDIR /lambda

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd

FROM alpine:3.16

COPY --from=build /lambda/main ./main

RUN apk --no-cache add ca-certificates

RUN chmod +x /main

ENTRYPOINT [ "/main" ]
