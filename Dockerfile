# build stage
FROM golang:1.15-alpine AS build-env
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app