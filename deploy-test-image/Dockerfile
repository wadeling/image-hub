FROM ubuntu:18.04
RUN apt-get update && apt-get install -y bash
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
