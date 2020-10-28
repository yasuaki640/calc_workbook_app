FROM golang:1.15-alpine3.12

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/bin \
    && cd /usr/bin \
    && ln -s migrate.linux-amd64 migrate

EXPOSE 8000

CMD ["fresh"]