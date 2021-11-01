#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)

deployPre(){
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    SERVICE="${PROJECT}"
    ssh -o stricthostkeychecking=no mqq@${HOST} -p 60022 "
        docker stop ${PROJECT}
        docker rm ${PROJECT}
        aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${IMAGE_REPO}
        docker run -d --restart=always --name=${PROJECT} --network host \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        npm run start
        docker container list
        "
}

deploy(){
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    SERVICE="${PROJECT}"
    ssh -o stricthostkeychecking=no mqq@${HOST} "
        docker stop ${PROJECT}
        docker rm ${PROJECT}
        aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${IMAGE_REPO}
        docker run -d --restart=always --name=${PROJECT} --network host \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        npm run start
        docker container list
        "
}

deploy_k8s() {
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    DEPLOYMENT="${PROJECT}"
    kubectl set image "deployment/${DEPLOYMENT}" ${DEPLOYMENT}="$IMAGE_REPO/$PROJECT:$VERSION" -n ficool
}


case ${CI_COMMIT_REF_NAME} in
  master|v*)
    ENV="prod"
    echo "${FICOOL_PROD}"
    if [[ "${FICOOL_PROD}" == "" ]];then
      echo "FICOOL_PROD not set"
      exit 1
    fi
    for HOST in ${FICOOL_PROD}
    do
      deploy ${ENV} "${HOST}"
    done
    ;;
  pre)
    ENV="pre"
    echo "${FICOOL_PRE}"
    if [[ "${FICOOL_PRE}" == "" ]];then
      echo "FICOOL_PRE not set"
      exit 1
    fi
    for HOST in ${FICOOL_PRE}
    do
      deployPre ${ENV} "${HOST}"
    done
    ;;
  *)
    ENV="dev"
    deploy_k8s ${ENV} "${HOST}"
    ;;
esac

