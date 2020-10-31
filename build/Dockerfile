# Builder
FROM golang:1.15.3-alpine3.12 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make docs
RUN make dependencies
RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/basesvc /app

CMD ["/app/basesvc", "api"]