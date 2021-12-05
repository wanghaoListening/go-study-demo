#!/bin/bash --login

#version 1.0.6 fix update scp to rsync
#owner= luoshiqi@bytedance.com
REMOTE_GO_PATH="/home/luoshiqi/go"
PROJECT_PATH="/src/git.byted.org/ee/people/survey/report"
#username@ip  ex:luoshiqi@10.xx.xx.xx
devbox=dev
PRODUCT="people"
SUBSYS="survey"
MODULE="report"
PORT=8888

REMOTE_PROJECT_PATH=${REMOTE_GO_PATH}/${PROJECT_PATH}
RUN_NAME=${PRODUCT}.${SUBSYS}.${MODULE}

echo kill start
ssh ${devbox} 'bash -s' < kill_dlv.sh ${RUN_NAME}
echo kill end

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o output/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
ssh ${devbox} mkdir -p ${REMOTE_PROJECT_PATH}/output
ssh ${devbox} mkdir -p ${REMOTE_PROJECT_PATH}/conf
echo 'rsync start'
rsync -r output/ ${devbox}:${REMOTE_PROJECT_PATH}/output
rsync -r conf/   ${devbox}:${REMOTE_PROJECT_PATH}/conf
echo 'rsync complete'
echo 'start debug.sh'
echo debug.sh ${REMOTE_PROJECT_PATH} ${REMOTE_GO_PATH} ${PRODUCT} ${SUBSYS} ${MODULE} ${PORT}
#> /dev/null
ssh ${devbox} 'bash -s' < debug.sh ${REMOTE_PROJECT_PATH} ${REMOTE_GO_PATH} ${PRODUCT} ${SUBSYS} ${MODULE} ${PORT} 2 > log.txt >&1 &
echo "exit debug.sh"

