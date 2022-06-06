FROM ubuntu:18.04
COPY [" .", "/home/ ."]
COPY [" ..", "/home/ .."]
COPY ["!", "/home/!"]
COPY ["@", "/home/@"]
COPY ["-", "/home/-"]
COPY ["#", "/home/#"]
# tricky way to add file '*' in macos,but not suitable for linux
#ADD ./"'*'" /home
COPY ?.sh /home/?.sh
COPY dp-test.sh /home/dp-test.sh
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN cd /home && \
ln -s dp-test.sh link1 
# link in diff layer
RUN cd /home && ln -s link1 link2 && \
ln -s linka linkb && \
ln -s linkb linka
# test overwrite
COPY dp-test2.sh /home/dp-test.sh

ENTRYPOINT ["entrypoint.sh"]
