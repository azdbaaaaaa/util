#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)
export VERSION
echo '$VERSION',$VERSION
IMAGE_REPO=${IMAGE_REPO}
export IMAGE_REPO
echo '$IMAGE_REPO',$IMAGE_REPO

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
        docker run -d --log-driver json-file --log-opt max-size=50m --log-opt max-file=10 --restart=always --name=${SERVICE} \
        --network host --env GIN_MODE=release\
        -v /log:/log \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        ${CMD} --config=/app/config/${PROJECT}-${ENV}.yaml
        docker container list
        "
}

deploy_k8s() {
    export ENV=$1
    export NAMESPACE=$2

    # 更新/新建 configmap
    kubectl create configmap "${PROJECT}-${CMD}" --from-file=${PROJECT}-${CMD}.yaml="config/${PROJECT}-${ENV}.yaml" -n "${NAMESPACE}" -o yaml --dry-run=client | kubectl replace -f - || \
    kubectl create configmap "${PROJECT}-${CMD}" --from-file=${PROJECT}-${CMD}.yaml="config/${PROJECT}-${ENV}.yaml" -n "${NAMESPACE}"

    export HTTP_PORT=`kubectl get configmap "${PROJECT}-${CMD}" -n "${NAMESPACE}" -o json | jq -r ".data.\"${PROJECT}-${CMD}.yaml\"" | yq e '.http.addr' - | sed 's/://g'`
    export GRPC_PORT=`kubectl get configmap "${PROJECT}-${CMD}" -n "${NAMESPACE}" -o json | jq -r ".data.\"${PROJECT}-${CMD}.yaml\"" | yq e '.grpc.addr' - | sed 's/://g'`
    export SCHEDULE=`kubectl get configmap "${PROJECT}-${CMD}" -n "${NAMESPACE}" -o json | jq -r ".data.\"${PROJECT}-${CMD}.yaml\"" | yq e ".cronjob.${CMD}.schedule" -`

    echo "开始发布，环境:",${ENV},'$NAMESPACE',$NAMESPACE,'$TYPE',$TYPE,'$INIT',$INIT,'$CMD',$CMD,'$HTTP_PORT',$HTTP_PORT,'$GRPC_PORT',$GRPC_PORT,'$SCHEDULE',$SCHEDULE

    if [[ $INIT == "" ]];then
      case "${TYPE}" in
      consumer|http|grpc)
        kubectl set image "deployment/${PROJECT}-${CMD}" "${PROJECT}-${CMD}"="$IMAGE_REPO/$PROJECT:$VERSION" -n "${NAMESPACE}"
        echo "deployment updated"
        exit 0
        echo "already return"
        ;;
      job)
        kubectl delete job "${PROJECT}-${CMD}" -n "${NAMESPACE}" || echo "job not found"
        echo "job deleted"
      esac
    fi

    echo "kustomize yaml"
    git clone https://github.com/azdbaaaaaa/util.git --depth=1
    echo "clone util finished"
    kustomize build util/scripts/k8s/app/${TYPE}/overlays/${ENV} > all.yaml
    echo "kustomize build finished"
    file=`cat all.yaml`
    export -p
    echo "export -p"
    printf "`export -p`\ncat << EOF\n$file\nEOF" | bash > all.yaml
    cat all.yaml
    echo "cat all.yaml"
    kubectl apply -f all.yaml
}

case ${CI_COMMIT_REF_NAME} in
  v*)
    ENV="prod"
    echo "${FICOOL_PROD}"
    if [[ "${FICOOL_PROD}" == "" ]];then
      echo "FICOOL_PROD not set"
      exit 1
    fi

    case "${TYPE}" in
      consumer|http)
        echo "开始部署k8s"
        NAMESPACE="prod-ficool"
        deploy_k8s ${ENV} ${NAMESPACE}
      ;;
      *)
        for HOST in ${FICOOL_PROD}
        do
          deploy ${ENV} "${HOST}"
        done
      ;;
    esac
    ;;
  pre)
    ENV="pre"
    echo "开始部署${FICOOL_PRE}"
    if [[ "${FICOOL_PRE}" == "" ]];then
      echo "FICOOL_PRE not set"
      exit 1
    fi

    case "${TYPE}" in
      cronjob|job|consumer)
        echo "cronjob|job|consumer do not need to deploy on host"
      ;;
      *)
        for HOST in ${FICOOL_PRE}
        do
          deployPre ${ENV} "${HOST}"
        done
      ;;
    esac
    echo "开始部署k8s"
    NAMESPACE="pre-ficool"
    deploy_k8s ${ENV} ${NAMESPACE}
    ;;
  master)
    echo "please set tags to publish!"
    exit 1
    ;;
  *)
    ENV="dev"
    NAMESPACE="ficool"
    deploy_k8s ${ENV} ${NAMESPACE}
    ;;
esac

