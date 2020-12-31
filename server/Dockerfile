FROM golang:1.15-alpine as base

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

RUN apk add --no-cache \
        alpine-sdk \
        git
RUN go mod download
RUN go get -u github.com/cosmtrek/air

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