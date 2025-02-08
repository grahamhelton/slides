#!/usr/bin/env bash
NOCOLOR=$(tput sgr0)
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
BLUE=$(tput setaf 4)
YELLOW=$(tput setaf 3)
TICK="$NOCOLOR[$GREEN+$NOCOLOR] "
TICK_ERROR="$NOCOLOR[$RED!$NOCOLOR] "

echo -n $TICK"Checking for etcd pod name in$BLUE kube-system$NOCOLOR namespace... "
ETCD_NAME=$(kubectl get pods -n kube-system | grep etcd | awk '{print $1}')
echo $YELLOW $ETCD_NAME
ETCD_INFO=$(kubectl describe pod -n kube-system $ETCD_NAME)
ETCD_CACERT=$(echo "$ETCD_INFO" | grep '\--trusted-ca-file'| cut -d"=" -f 2)
ETCD_SERVERCERT=$(echo "$ETCD_INFO" | grep '\--cert-file' | cut -d"=" -f 2)
ETCD_KEY=$(echo "$ETCD_INFO" | grep '\--key-file' | cut -d"=" -f 2)

echo $TICK"Attempting to save etcd databse snapshot to $BLUE/tmp/etcd-loot.db"$NOCOLOR
ETCDCTL_API=3 etcdctl --cacert=$ETCD_CACERT --cert=$ETCD_SERVERCERT --key=$ETCD_KEY snapshot save /tmp/etcd-loot.db
if [ $? -eq 0 ];then
        echo $TICK"Etcd snapshot success, stored in $BLUE/tmp/etcd-loot.db!"$NOCOLOR
else
        echo $TICK_ERROR$RED"Failed to take snapshot of etcd database!"$NOCOLOR
fi
