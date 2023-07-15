FROM alpine:3.14
RUN apk add bash 
COPY test /usr/local/bin/test
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
