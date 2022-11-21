FROM debian:11.3
ADD linux-libc-dev_5.10.120-1+grm3u2_amd64.deb /
RUN dpkg -i linux-libc-dev_5.10.120-1+grm3u2_amd64.deb
RUN apt update && apt upgrade linux-libc-dev
