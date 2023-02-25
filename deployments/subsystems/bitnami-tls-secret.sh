# Requires the following variables defined:
# EXTERNAL_NAMESPACES, EXTERNAL_SECRET, TLS_SECRET

# EXTERNAL_NAMESPACES -> namespaces separated by a space, e.g. "a b c"

# Warning: depends on helm-deployment.sh

check_variables \
  EXTERNAL_NAMESPACES \
  EXTERNAL_SECRET \
  TLS_SECRET

# Changes secret to its final result
SECRET_FILE=$(kubectl get secret $TLS_SECRET -n $NAMESPACE -o json \
  | jq 'del(.metadata["namespace","creationTimestamp","resourceVersion","selfLink","uid"])|.metadata.name="'"$EXTERNAL_SECRET"'"')

for ns in $EXTERNAL_NAMESPACES; do
echo $SECRET_FILE | kubectl -n $ns apply -f -
done
