FROM golang:1.15.3-alpine3.12

RUN apk update && apk upgrade && \
    apk --update add curl make git

WORKDIR /app

COPY . .

RUN make docs
RUN make dependencies

CMD ["air", "-c", "/app/.air.toml"]