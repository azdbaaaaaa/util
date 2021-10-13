#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)

deployPre(){
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    if [[ ${CMD} == "serve" ]]
    then
      SERVICE="${PROJECT}"
    else
      SERVICE="${PROJECT}-${CMD}"
    fi
    ssh -o stricthostkeychecking=no mqq@${HOST} -p 60022 "
        docker stop ${SERVICE}
        docker rm ${SERVICE}
        aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${IMAGE_REPO}
        docker run -d --restart=always --name=${SERVICE} --network host \
        -v /usr/local/app/tars/app_log/LightHouse:/usr/local/app/tars/app_log/LightHouse \
        -v /log:/log \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        ${CMD} --config=/app/config/${PROJECT}-${ENV}.yaml
        docker container list
        "
}

deploy(){
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    if [[ ${CMD} == "serve" ]]
    then
      SERVICE="${PROJECT}"
    else
      SERVICE="${PROJECT}-${CMD}"
    fi
    ssh -o stricthostkeychecking=no mqq@${HOST} "
        docker stop ${SERVICE}
        docker rm ${SERVICE}
        aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${IMAGE_REPO}
        docker run -d --restart=always --name=${SERVICE} --network host \
        -v /log:/log \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        ${CMD} --config=/app/config/${PROJECT}-${ENV}.yaml
        docker container list
        "
}

deploy_k8s() {
    ENV=$1
    HOST=$2
    echo "开始发布，主机ip为:${HOST}"
    echo ${PROJECT},${IMAGE_REPO},${VERSION},${ENV},${HOST},${CMD}
    if [[ ${CMD} == "serve" ]]
    then
      DEPLOYMENT="${PROJECT}"
    else
      DEPLOYMENT="${PROJECT}-${CMD}"
    fi

    kubectl create configmap "${PROJECT}" --from-file=${PROJECT}.yaml="config/${PROJECT}-${ENV}.yaml" -n ficool -o yaml --dry-run=client | kubectl replace -f -
    kubectl set image "deployment/${DEPLOYMENT}" ${DEPLOYMENT}="$IMAGE_REPO/$PROJECT:$VERSION" -n ficool
}


case ${CI_COMMIT_REF_NAME} in
  v*)
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
  master)
    echo "please set tags to publish!"
    exit 1
    ;;
  *)
    ENV="dev"
    deploy_k8s ${ENV} "${HOST}"
    ;;
esac

