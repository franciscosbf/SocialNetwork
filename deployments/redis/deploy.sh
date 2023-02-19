#!/usr/bin/env sh

REPO_NAME="bitnami"
REPO_URL="https://charts.bitnami.com/bitnami"
RELEASE_NAME="redis-cluster"
VALUES_FILE="values.yaml"
NAMESPACE_FILE="redis-namespace.yaml"

source ../common/helm-deployment.sh
