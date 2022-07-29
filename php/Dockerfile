FROM debian:11

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    apt-get update && apt-get install php7.4-common=7.4.28-1+deb11u1 -y && \
    apt-get install php7.4-pgsql=7.4.28-1+deb11u1 -y
