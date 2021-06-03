FROM golang:1.16-alpine as build

WORKDIR /srv/app
COPY . .

RUN make build

FROM alpine:3
LABEL "org.opencontainers.image.source"="https://github.com/JDR-ynovant/api"

WORKDIR /srv/app
COPY --from=build /srv/app/candy-fight .

EXPOSE 3000
CMD ["/srv/app/api", "serve"]
