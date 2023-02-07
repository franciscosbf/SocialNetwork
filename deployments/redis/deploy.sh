#!/usr/bin/env sh

REPO_NAME="bitnami"
REPO_URL="https://charts.bitnami.com/bitnami"
RELEASE_NAME="redis"
VALUES_FILE="values.yaml"

source ../common/helm-deployment.sh
