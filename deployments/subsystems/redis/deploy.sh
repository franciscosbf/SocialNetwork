#!/usr/bin/env sh

REPO_NAME="bitnami"
REPO_URL="https://charts.bitnami.com/bitnami"
RELEASE_NAME="redis-cluster"
CHART_NAME="redis-cluster"
VALUES_FILE="values.yaml"
NAMESPACE_FILE="redis-namespace.yaml"

source ../helm-deployment.sh

EXTERNAL_NAMESPACES="accounts"
EXTERNAL_SECRET="redis-tls"
TLS_SECRET="redis-cluster-crt"

source ../bitnami-tls-secret.sh

