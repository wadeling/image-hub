FROM alpine:3.14
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN apk add bash && apk add musl musl-utils pkgconf
COPY entrypoint.sh /root/entrypoint.sh
RUN apk del pkgconf
ENTRYPOINT ["entrypoint.sh"]
