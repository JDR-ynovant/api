FROM golang:1.16-alpine as build

WORKDIR /srv/app
COPY . .

RUN go mod download && \
    go build

FROM alpine
LABEL "org.opencontainers.image.source"="https://github.com/JDR-ynovant/api"

WORKDIR /srv/app
COPY --from=build /srv/app/api

EXPOSE 3000
CMD ["/srv/app/api", "serve"]
