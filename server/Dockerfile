FROM golang:1.15-alpine as base

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

RUN apk add --no-cache \
        alpine-sdk \
        git
RUN go mod download
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv

# create link for migration tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/bin \
    && cd /usr/bin \
    && ln -s migrate.linux-amd64 migrate

COPY . .

FROM golang:1.15-alpine as builder

WORKDIR /app/go/builder

COPY --from=base /app/go/base /app/go/builder

RUN CGO_ENABLED=0 go build main.go


FROM alpine as production

WORKDIR /app/go/src

COPY --from=builder /app/go/builder/main /app/go/src/main

EXPOSE 8000

CMD ["./main"]