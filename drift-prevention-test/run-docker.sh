#!/bin/bash
#docker run -d --name dp-test wade23/deploy:dp-test

docker run -d --name dp-test -v /proc:/host/proc wade23/deploy:dp-test
