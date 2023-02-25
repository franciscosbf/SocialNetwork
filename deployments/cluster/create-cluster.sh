#!/usr/bin/env sh

# Credits: https://kind.sigs.k8s.io/docs/user/loadbalancer/

panic() {
  echo "cluster creation error: $@"; exit 1
}

print() {
  echo "cluster creation info: $@"
}

source ./cluster-name.sh

CLI_CONF="$HOME/.kube/config"
if ! [ -z "$1" ]; then
  CLI_CONF=$1
fi

install() {
  HAS_KIND=$(command -v kind)
  if ! [ "$HAS_KIND" ]; then
    panic "kind is missing"
  fi

  CLUSTER_CONFIG="kind-config.yaml"

  if ! [ -f $CLUSTER_CONFIG ]; then
    panic "config file $CLUSTER_CONFIG is missing"
  fi

  CLUSTER_NAME=$1
  KUBECTL_CONF=$2

  # Create cluster and store client config
  kind create cluster --config $CLUSTER_CONFIG --name $CLUSTER_NAME && \
    # Fetches kube config and stores it
    kind get kubeconfig --name $CLUSTER_NAME > $KUBECTL_CONF && \
    print "kubectl config stored in $KUBECTL_CONF" && \

    # Apply MetaLB manifest
    kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml && \

    # Wait until pods are ready
    kubectl wait --namespace metallb-system \
        --for=condition=ready pod \
        --selector=app=metallb \
        --timeout=90s && \

    kubectl apply -f - << EOF
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: metallb
  namespace: metallb-system
spec:
  addresses:
  - $(docker network inspect -f '{{(index .IPAM.Config 0).Subnet}}' kind) # IPv6
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: metallb
  namespace: metallb-system
EOF
}

read -p "Do you need sudo? [Y/N] " yn
case $yn in
    [Yy]* ) DECLARATIONS=$(declare -f install panic print); \
              sudo sh -c "$DECLARATIONS; install $CLUSTER $CLI_CONF"; \
              break;;
    * ) install $CLI_CONF $CLUSTER;;
esac
