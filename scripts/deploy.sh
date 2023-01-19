#!/usr/bin/env bash
set -e

PROJECT="${CI_PROJECT_NAME}"
IMAGE_REPO="991619025488.dkr.ecr.eu-west-1.amazonaws.com"
echo "当前项目为:${PROJECT}"


# shellcheck disable=SC2006
DATE=`date '+%Y%m%d-%H%M%S'`
# shellcheck disable=SC2006
VERSION=`git describe --tags --always`
VERSION="${DATE}-${VERSION}"


deploy(){
    PROJECT=$1
    IMAGE_REPO=$2
    VERSION=$3
    ENV=$4
    HOST=$5
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST}
    ssh root@${HOST} -p 60022 "
        docker stop ${PROJECT}
        docker rm ${PROJECT}
        aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${IMAGE_REPO}
        docker run -d --restart=always --name=${PROJECT} --network host \
        -v /usr/local/app/tars/app_log/LightHouse:/usr/local/app/tars/app_log/LightHouse \
        -v /etc/localtime:/etc/localtime:ro \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        serve --config=/app/config/${PROJECT}-${ENV}.yaml
        docker container list
        "
}

# shellcheck disable=SC2006
PWD=`pwd`
echo "CI_COMMIT_REF_NAME:${CI_COMMIT_REF_NAME}"
echo "CI_DEFAULT_BRANCH:${CI_DEFAULT_BRANCH}"
echo "CI_COMMIT_BRANCH:${CI_COMMIT_BRANCH}"
echo "CI_PROJECT_NAME:${CI_PROJECT_NAME},CI_PROJECT_ID:${CI_PROJECT_ID},CI_PROJECT_TITLE:${CI_PROJECT_TITLE}"
echo "当前使用CI_COMMIT_REF_NAME:${CI_COMMIT_REF_NAME}"
echo "当前目录为:${PWD}"
case ${CI_COMMIT_REF_NAME} in
  pre|pre*)
    ENV="pre"
    echo "开始构建任务，版本为:${VERSION},环境：${ENV}"
    echo "开始打包"
    make build PROJECT="${PROJECT}" IMAGE_REPO="${IMAGE_REPO}" VERSION="${VERSION}"
    echo "完成打包"
    for HOST in ${FICOOL_PRE}
    do
      echo "开始发布，主机ip为:${HOST}"
      # shellcheck disable=SC2086
      deploy ${PROJECT} ${IMAGE_REPO} ${VERSION} ${ENV} ${HOST}
    done
    ;;
  master|v*)
    ENV="prod"
    echo "开始构建任务，版本为:${VERSION},环境：${ENV}"
    echo "开始打包"
    make build PROJECT="${PROJECT}" IMAGE_REPO="${IMAGE_REPO}" VERSION="${VERSION}"
    echo "完成打包"
    for HOST in ${FICOOL_PROD}
    do
      echo "开始发布，主机ip为:${HOST}"
      # shellcheck disable=SC2086
      deploy ${PROJECT} ${IMAGE_REPO} ${VERSION} ${ENV} ${HOST}
    done
    ;;
  *)
    ENV="dev"
    echo "开始构建任务，版本为:${VERSION},环境：${ENV}"
    echo "开始打包"
    make build PROJECT="${PROJECT}" IMAGE_REPO="${IMAGE_REPO}" VERSION="${VERSION}"
    echo "完成打包"
    # shellcheck disable=SC2086
    kubectl create configmap ${PROJECT} --from-file=${PROJECT}.yaml=config/${PROJECT}-${ENV}.yaml -n ficool -o yaml --dry-run=client | kubectl replace -f -
    # shellcheck disable=SC2086
    kubectl set image deployment/${PROJECT} ${PROJECT}=${IMAGE_REPO}/${PROJECT}:${VERSION} -n ficool
    for HOST in ${FICOOL_DEV}
    do
      # shellcheck disable=SC2086
      deploy ${PROJECT} ${IMAGE_REPO} ${VERSION} ${ENV} ${HOST}
    done
    ;;
esac

