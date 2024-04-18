# Notes

Add new context for the simple-bank cluster to ./kube/config file:
`aws eks update-kubeconfig --name simple-bank --region ap-southeast-1`

Switch context:
`kubectl config use-context arn:aws:eks:ap-southeast-1:929684159033:cluster/simple-bank`

[Amazon EKS cluster accesss](https://repost.aws/knowledge-center/amazon-eks-cluster-access)

- `aws sts get-caller-identity`
- `cat ~/.aws/credentials`
- `vi ~/.aws/credentials` -> change access and secret key from AWS security credentials
- `kubectl get pods`
- `kubectl cluster-info`

Give user access to github-ci to cluster, since we want github to automatically deploy the app, whenever we push new changes to the main branch:

- `export AWS_PROFILE=github` -> give aws access to github
- create `aws-auth.yaml`
- `export AWS_PROFILE=default`
- `kubectl apply -f eks/aws-auth.yaml` -> give access to user who's not the create of the cluster

[k9s](https://k9scli.io/)
