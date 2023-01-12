#!/bin/bash

docker run -d --name test-http-server -p 8082:8082 wade23/dev:http-server
