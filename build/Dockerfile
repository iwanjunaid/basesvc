# Builder
FROM golang:1.15.3-alpine3.12 as builder

RUN apk update && apk upgrade && \
    apk --update add git make curl

WORKDIR /app

COPY . .
COPY config/.air.toml .air.toml

RUN make dependencies

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && ./bin/air /bin/air

EXPOSE 9090

CMD ["air", "-c", ".air.toml"]