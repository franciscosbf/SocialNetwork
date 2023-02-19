# Requires the following variables defined:
# REPO_NAME, REPO_URL, RELEASE_NAME, VALUES_FILE

# Echoes something and exists with 1
panic() {
  echo "error: $@"; exit 1
}

# Echoes something
print() {
  echo "info: $@"
}

# Checks if a given set of variables is defined
check_variables() {
  for var in $@; do
    if [ -z ${!var+x} ]; then
      panic "missing variable $var"
    fi
  done
}

# Checks if a given set of commands exist
check_commands() {
  for comm in $@; do
    HAS_COMMAND=$(command -v $comm)
    if ! [ "$HAS_COMMAND" ]; then
      panic "$comm command is required"
    fi
  done
}

# Checks if a given set of files exist
check_files () {
  for file in $@; do
    if ! [ -f $file ]; then
      panic "$file is missing"
    fi
  done
}

check_variables \
  REPO_NAME \
  REPO_URL \
  RELEASE_NAME \
  VALUES_FILE \
  NAMESPACE_FILE

check_commands \
  helm \
  jq \
  yq \
  kubectl

check_files \
 $VALUES_FILE \
 $NAMESPACE_FILE

# Checks if the repo was already added. If not, adds it
HAS_REPO=$(helm repo list -o json | jq ".[].name==\"$REPO_NAME\"")
if [ "$HAS_REPO" = "false" ]; then
  print "adding $REPO_URL as $REPO_NAME"
  helm repo add $REPO_NAME $REPO_URL
fi

# Checks if release is already installed
HAS_RELEASE=$(helm list -l name=$RELEASE_NAME -o json)
if [ "$HAS_RELEASE" != "[]" ]; then
  panic "$RELEASE_NAME release is already installed"
fi

# Checks if contains the name label
HAS_LABEL=$(cat redis-namespace.yaml | yq -M -r ".metadata.labels.name")
if [ "$HAS_LABEL" = "null" ]; then
  panic "missing .metadata.labels.name in namespace file $NAMESPACE_FILE"
fi
NAMESPACE=$HAS_LABEL

# Applies required namespace
kubectl apply -f $NAMESPACE_FILE

# Installs the release
print "installing as $RELEASE_NAME"
helm install $RELEASE_NAME -n $NAMESPACE -f $VALUES_FILE $REPO_NAME/$RELEASE_NAME
