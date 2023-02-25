#!/usr/bin/env sh

source ./cluster-name.sh

delete() {
  kind delete cluster -n $1
}

read -p "Do you need sudo? [Y/N] " yn
case $yn in
    [Yy]* ) DECLARATIONS=$(declare -f delete); \
              sudo sh -c "$DECLARATIONS; delete $CLUSTER"; \
              break;;
    * ) delete $CLUSTER;;
esac
