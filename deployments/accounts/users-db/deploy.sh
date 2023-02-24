#!/usr/bin/env sh

REPO_NAME="bitnami"
REPO_URL="https://charts.bitnami.com/bitnami"
RELEASE_NAME="accounts-db"
CHART_NAME="postgresql"
VALUES_FILE="values.yaml"
NAMESPACE_FILE="../accounts-namespace.yaml"

source ../../common/helm-deployment.sh
