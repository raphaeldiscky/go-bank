# Notes

## Connect k8s cluster on AWS EKS

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

## Deploy web app to k8s cluster on AWS EKS

- create `deployment.yaml`
- `kubectl apply -f eks/deployment.yaml`
- use k9s to easy UI

Route traffic from the outside world to the pod:

- create `service.yaml` with type `LoadBalancer` to add external-ip / expose service to the outside world [Service](https://kubernetes.io/docs/concepts/services-networking/service/)

- `kubectl apply -f eks/service.yaml`
- check connection `nslookup a6e79e3e6b14340cda36500aa21972ed-709104643.ap-southeast-1.elb.amazonaws.com`
- now to test in postman `http://a6e79e3e6b14340cda36500aa21972ed-709104643.ap-southeast-1.elb.amazonaws.com/users/login`

## Register a domain and setup A-record

- setup in AWS Route53
- in Postman change this to `http://{domain}/users/login`

## User Ingress to route traffics to different services in k8s

[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)

- change `service.yaml` type to `ClusterIP` since we don't want to expose service to outside world anymore
- create `ingress.yaml` -> doesn't have external IP -> to make it work use [Ingress Controllers](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/)
- add class `IngressClass` to `ingress.yaml`
- `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/aws/deploy.yaml`
- copy ingress Address
- edit route traffic to with previous address in AWS Route53

## Automatic issue TLS certificates in k8s -> Let's Encrypt

[cert-manager](https://cert-manager.io/)

- use Let's Encrypt to deploy cert-manager and configure it to get certificates for the NGIX Ingress controller
- the best way is to use DNS-01 challenge instead of HTTP-01 challenge since it can manage wildcard certificates. Use it only if the DNS provider has an API to automate the record updates
- for simpler way -> we use HTTP-01 challenge
- installation [cert-manager](https://cert-manager.io/docs/installation/kubectl/)
- config issuers to the cluster -> [ACME Issuer](https://cert-manager.io/docs/configuration/acme/)
- create `issuer.yaml`
- `kubectl apply -f eks/issuer.yaml`
- attach issuer to ingress in `ingress.yaml`
- `kubectl apply -f eks/ingress.yaml`
- in Postman add https to `https://{domain}/users/login`

## Automatic deploy to k8s with Github Action

- [kubectl github action](https://github.com/marketplace/actions/kubectl-tool-installer)
- edit `deploy.yaml`
- we want -> whenever we push new changes to the main branch, it will build and tag the latest version image and the image should be redeployed to the k8s cluster
