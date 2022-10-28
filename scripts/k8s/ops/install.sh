
# https://docs.aws.amazon.com/zh_cn/eks/latest/userguide/efs-csi.html
cluster="prod-ficool"
namespace="kube-system"
policy_arn="arn:aws:iam::991619025488:policy/AmazonEKS_EFS_CSI_Driver_Policy"
region="eu-west-1"
eksctl utils associate-iam-oidc-provider --region=eu-west-1 --cluster=prod-ficool --approve
eksctl create iamserviceaccount \
    --cluster ${cluster} \
    --namespace ${namespace} \
    --name efs-csi-controller-sa \
    --attach-policy-arn ${policy_arn} \
    --approve \
    --region ${region}
helm repo add aws-efs-csi-driver https://kubernetes-sigs.github.io/aws-efs-csi-driver/
helm repo update
helm upgrade -i aws-efs-csi-driver aws-efs-csi-driver/aws-efs-csi-driver \
    --namespace ${namespace} \
    --set image.repository=602401143452.dkr.ecr.eu-west-1.amazonaws.com/eks/aws-efs-csi-driver \
    --set controller.serviceAccount.create=false \
    --set controller.serviceAccount.name=efs-csi-controller-sa

kubectl apply -f storage/storageclass.yaml
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml
