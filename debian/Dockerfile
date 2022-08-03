FROM debian:stretch
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    apt-get update && apt-get install tzdata=2020a-0+deb9u1 -y --allow-downgrades
