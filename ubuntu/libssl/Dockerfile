FROM ubuntu:jammy
RUN sed -i 's!archive.ubuntu.com/!mirrors.aliyun.com/!g' /etc/apt/sources.list \
    && sed -i 's!security.ubuntu.com/!mirrors.aliyun.com/!g' /etc/apt/sources.list \
    && apt-get update \
    && apt-get full-upgrade -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
