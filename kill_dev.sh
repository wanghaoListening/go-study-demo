#!/usr/bin/env bash

echo $1

cnt=0
cnt=$(ps -ef | grep $1 | wc -l)

if [ $cnt -gt 0 ]
then
    ps -ef | grep $1 | awk '{print $2}'| xargs kill -9 &
fi

sleep 3
echo 'kill dlv process success'
