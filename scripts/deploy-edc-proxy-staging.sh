set -e

export IMAGE_TAG="${1:-v0.0.1}"

echo "Deploying K8S resources..."
export KUBECONFIG="/cygdrive/c/Users/dkaragkounis/Desktop/Omega-X/omega-x.yaml"
cat deployments/edc-proxy-staging.yaml| envsubst | kubectl -n marketplace apply -f -
kubectl -n marketplace wait --for=condition="Available" --timeout=60s deployment/edc-proxy-deployment-staging
echo "Deployment is now available."
kubectl -n marketplace get pods
echo "K8S resources deployment completed."
