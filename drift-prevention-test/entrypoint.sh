#!/bin/bash

while true
do
	echo "hello world\n"
	pid=`cut -f 4 -d ' ' /host/proc/self/stat`
	echo "my pid:$pid\n"
	res=`ls -l /host/proc/self`
	echo "res: $res"
	sleep 10
done
