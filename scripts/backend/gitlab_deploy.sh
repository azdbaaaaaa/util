#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)
export VERSION
IMAGE_REPO=${IMAGE_REPO}
export IMAGE_REPO


yaml_deployment(){
  export PORT=$PORT
  wget https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/${ENV}/${TYPE}/deployment.template.yaml
  cp -a deployment.template.yaml deployment.yaml
  file=`cat deployment.yaml`
  printf "`export -p`\ncat << EOF\n$file\nEOF" | bash > deployment.yaml
}

yaml_service(){
  export PORT=$PORT
  wget https://raw.githubusercontent.com/azdbaaaaaa/util/master/scripts/k8s/${ENV}/${TYPE}/service.template.yaml
  cp -a service.template.yaml service.yaml
  file=`cat service.yaml`
  printf "`export -p`\ncat << EOF\n$file\nEOF" | bash > service.yaml
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
        docker run -d --log-driver json-file --log-opt max-size=50m --log-opt max-file=10 --restart=always --name=${SERVICE} --network host \
        -v /log:/log \
        ${IMAGE_REPO}/${PROJECT}:${VERSION} \
        ${CMD} --config=/app/config/${PROJECT}-${ENV}.yaml
        docker container list
        "
}

deploy_k8s() {
    export ENV=$1
    export NAMESPACE=$2

    echo "开始发布，环境:${ENV}"
    if [[ ${CMD} == "serve" ]]
    then
      export DEPLOYMENT="${PROJECT}"
      export SERVICE="${PROJECT}"
      if [[ ${PROJECT} == 'go-api-service' ]];then
        export TYPE="http"
      else
        export TYPE="grpc"
      fi
    else
      export DEPLOYMENT="${PROJECT}-${CMD}"
      export SERVICE="${PROJECT}-${CMD}"
      export TYPE="job"
    fi
    ## k8s configmap
    kubectl create configmap "${DEPLOYMENT}" --from-file=${DEPLOYMENT}.yaml="config/${PROJECT}-${ENV}.yaml" -n "${NAMESPACE}" ||
    kubectl create configmap "${DEPLOYMENT}" --from-file=${DEPLOYMENT}.yaml="config/${PROJECT}-${ENV}.yaml" -n "${NAMESPACE}" -o yaml --dry-run=client | kubectl replace -f -

    case "${TYPE}" in
    http)
      PORT=`kubectl get configmap ${DEPLOYMENT} -n "${NAMESPACE}" -o json | jq -r ".data.\"$DEPLOYMENT.yaml\"" | yq e '.http.addr' - | sed 's/://g'`
      export PORT=$PORT
      ;;
    grpc)
      PORT=`kubectl get configmap ${DEPLOYMENT} -n "${NAMESPACE}" -o json | jq -r ".data.\"$DEPLOYMENT.yaml\"" | yq e '.grpc.addr' - | sed 's/://g'`
      export PORT=$PORT

      ## k8s deployment
#      is_deploy=`kubectl get deploy -n "${NAMESPACE}" | grep "${DEPLOYMENT}" | wc -l`
#      if [[ $is_deploy -gt 0 ]];then
#        kubectl set image "deployment/${DEPLOYMENT}" ${DEPLOYMENT}="$IMAGE_REPO/$PROJECT:$VERSION" -n "${NAMESPACE}"
#      else
#        yaml_deployment
#        kubectl apply -f deployment.yaml
#      fi
#
#      ## k8s service
#      is_service=`kubectl get svc -n "${NAMESPACE}" | grep "${SERVICE}" | wc -l`
#      if [[ is_service -eq 0 ]];then
#        yaml_service
#        kubectl apply -f service.yaml
#      fi
      ;;
    job)
      ## k8s deployment
#      is_deploy=`kubectl get deploy -n "${NAMESPACE}" | grep "${DEPLOYMENT}" | wc -l`
#      if [[ $is_deploy -gt 0 ]];then
#        kubectl set image "deployment/${DEPLOYMENT}" ${DEPLOYMENT}="$IMAGE_REPO/$PROJECT:$VERSION" -n "${NAMESPACE}"
#      else
#        yaml_deployment
#        kubectl apply -f deployment.yaml
#      fi
      ;;
    esac

        git clone https://github.com/azdbaaaaaa/util.git
        kustomize build util/scripts/k8s/${TYPE}/overlays/${ENV} > all.yaml
        file=`cat all.yaml`
        printf "`export -p`\ncat << EOF\n$file\nEOF" | bash > all.yaml
        cat all.yaml
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
    for HOST in ${FICOOL_PROD}
    do
      deploy ${ENV} "${HOST}"
    done
    ;;
  pre)
    ENV="pre"
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

