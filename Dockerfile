# syntax=docker/dockerfile:1.4
FROM golang:1.21.1-alpine3.18 AS builder

LABEL maintainer="Mario Valentino <mariovalentino.lumbantobing@gmail.com>"
LABEL description="simple rest api with gin and gorm"

RUN apk --no-cache add ca-certificates

WORKDIR "/build"

COPY . .

RUN --mount=type=ssh \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build,id=gin-user-tasks \
go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o /build/gin-user-tasks .

FROM alpine:3.18

WORKDIR "/app"

RUN apk --no-cache add curl ca-certificates
RUN mkdir -p /app/logs

COPY --from=builder /build/.env /app

COPY --from=builder /build/gin-user-tasks /app

EXPOSE 8080

CMD [ "./gin-user-tasks" ]
