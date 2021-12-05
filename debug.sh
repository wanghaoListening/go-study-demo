#!/bin/bash --login


REMOTE_PROJECT_PATH=$1
REMOTE_GO_PATH=$2
PRODUCT=$3
SUBSYS=$4
MODULE=$5
PORT=$6

SVC_NAME=${PRODUCT}.${SUBSYS}.${MODULE}
EXEC_FILE=${REMOTE_PROJECT_PATH}/output/${SVC_NAME}
CONF_FILE=${REMOTE_PROJECT_PATH}/conf/${PRODUCT}_${SUBSYS}_${MODULE}.yml
dlv_cmd=${REMOTE_GO_PATH}/bin/dlv
if [ ! -f ${dlv_cmd} ];then
    echo "dlv 文件不存在 请先安装dlv "
    exit -1
fi

chmod 777 ${EXEC_FILE}
cmd=${dlv_cmd}" --headless --listen=:2345 --api-version=2 exec "${EXEC_FILE}" -- -conf="${CONF_FILE}" -log=. -svc="${SVC_NAME}"  -port="${PORT}
echo ${cmd}
${cmd}
#nohup ${cmd} &
#ps aux | grep ${SVC_NAME}
echo "dlv debug start ...."
exit 0
