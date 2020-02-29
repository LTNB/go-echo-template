FROM alpine:3.8
MAINTAINER Lamtnb <baolam0307@gmail.com>

RUN mkdir /app
COPY server /app/server
COPY configs /app/config
COPY views /app/views
COPY public /app/public
RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /app/server \
    && rm -rf /var/cache/apk/*

WORKDIR /app
ENTRYPOINT /app/server
