FROM ubuntu:22.04
RUN apt update && apt install -y curl
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
