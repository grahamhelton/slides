#! /bin/bash

# NOTE: To make use of a wider color pallet set the TERM enviornment variable TERM=xterm-256color

# Colors
BOLD=$(tput bold)
NOCOLOR=$(tput sgr0)
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
BLUE=$(tput setaf 4)
PURPLE=$(tput setaf 5)
CYAN=$(tput setaf 6)
WHITE=$(tput setaf 7)
BLACK=$(tput setaf 8)
BG_YELLOW=$(tput setab 3)
BOLD_RED=($BOLD$RED)

# Formatting
DIV="$BLACK---------------------------------------------------------------------$NOCOLOR"
TICK="$NOCOLOR[$GREEN+$NOCOLOR] "
TICK_MOVE="$NOCOLOR[$GREEN~>$NOCOLOR]"
TICK_BACKUP="$NOCOLOR[$GREEN<~$NOCOLOR] "
TICK_INPUT="$NOCOLOR[$YELLOW!$NOCOLOR] "
TICK_ERROR="$NOCOLOR[$RED!$NOCOLOR] "
TICK_INFO="$NOCOLOR[$YELLOW-$NOCOLOR] "


# Highlight
CRITICAL=($BG_YELLOW$BOLD_RED)

clear

spin() {
    local pid=$!
    local spin='|/-\'
    local i=0
    while kill -0 $pid 2>/dev/null; do
        i=$(((i + 1) % 4))
        printf "\r${spin:$i:1}"
        sleep .1
    done
    printf "  \r"
}

YAML="./create_pod.yaml"


echo $TICK"Creating resources for demo..."

kubectl apply -f $YAML
clear

echo $TICK"Dropping you into a pod in$BLUE dmz$NOCOLOR namespace"
echo $DIV
echo "Your goal is to get access to other areas of the cluster" 

sleep 10 & spin
echo $TICK"Waiting$BLUE 10 seconds$NOCOLOR for pod to start..."

kubectl exec -it pod-creator -n dmz -- bash

echo $TICK"Run$GREEN kubectl delete -f $YAML --force $NOCOLOR to destroy resources"
