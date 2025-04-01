set -e

export IMAGE_TAG="${1:-v0.0.1}"

echo "Deploying K8S resources..."
cat deployments/edc-proxy.yaml| envsubst | kubectl -n marketplace apply -f -
# kubectl -n marketplace apply -f deployments/edc-proxy.yaml
kubectl -n marketplace wait --for=condition="Available" --timeout=60s deployment/edc-proxy-deployment
echo "Deployment is now available."
kubectl -n marketplace get pods
echo "K8S resources deployment completed."
