## kind搭建集群
### install kind
```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
```
### install kubectl
```bash
curl -LO "https://dl.k8s.io/release/v1.22.7/bin/linux/amd64/kubectl"
docker pull kindest/node:v1.22.7
kind create cluster --image=kindest/node:v1.22.7 --name test
kubectl cluster-info --context kind-test
kubectl cluster-info # 查看k8s 集群信息
kind get clusters
kind delete cluster --name test
```