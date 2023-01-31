#! /usr/bin/env bash

APP_NAME=Kyanite
ARTIFACT_NAME=$(echo ${APP_NAME} | tr A-Z a-z)

MAJOR_VERSION=0
MINOR_VERSION=2
BUILD_VERSION=1

# 镜像tag
IMG_BASE_TAG=v${MAJOR_VERSION}.${MINOR_VERSION}.${BUILD_VERSION}

build_dev(){
    make
    # build debian based image
    nerdctl build -t ${ARTIFACT_NAME}:${IMG_BASE_TAG}-debian11-dev.$(date +%y%m%d) -f Dockerfile-debian11 .
    # build distroless based image
    nerdctl build -t ${ARTIFACT_NAME}:${IMG_BASE_TAG}-dev.$(date +%y%m%d) -f Dockerfile-distroless .
}

build_stable() {
    # build debian based image
    nerdctl build -t ${ARTIFACT_NAME}:${IMG_BASE_TAG}-debian11 -f Dockerfile-debian11 .
    # build distroless based image
    nerdctl build -t ${ARTIFACT_NAME}:${IMG_BASE_TAG} -f Dockerfile-distroless .
}

# main
if [[ ${#} -ne 1 ]];then
    echo "invalid args"
    exit 1
fi

case ${1} in
    "dev") build_dev
    ;;
    "stable") build_stable
    ;;
    *) echo "invalid option '${1}'"
    exit 1
    ;;
esac