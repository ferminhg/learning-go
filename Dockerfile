FROM golang:alpine AS build

RUN apk add --update make

ENV APP_HOME /app

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY . .
RUN CGO_ENABLED=0 go build -o /app/build/server /app/cmd/api/server.go

# Building image with the binary
FROM scratch
COPY --from=build /app/build/server /app/build/server
ENTRYPOINT ["/app/build/server"]
