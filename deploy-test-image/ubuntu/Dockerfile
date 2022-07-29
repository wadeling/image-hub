FROM ubuntu:18.04
RUN apt-get update && apt-get install -y bash && apt-get install -y python3
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
