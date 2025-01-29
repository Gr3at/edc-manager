# make sure required secrets have been created
# kubectl -n marketplace create secret generic edc-proxy-secrets --from-env-file=.env.prod --dry-run=client -o yaml
# kubectl -n marketplace create secret generic marketplace-registry-pull-secrets --from-file=.dockerconfigjson=.docker/config.json --type=kubernetes.io/dockerconfigjson --dry-run=client -o yaml

set -e

export IMAGE_TAG="${1:-v0.0.1}"

echo "Automigrate database tables schema..."
export KUBECONFIG="/cygdrive/c/Users/dkaragkounis/Desktop/Omega-X/omega-x.yaml"
# kubectl -n marketplace apply -f deployments/edc-proxy-migration-job.yaml
cat deployments/edc-proxy-migration-job.yaml| envsubst | kubectl -n marketplace apply -f -
kubectl -n marketplace wait --for=condition="Complete" --timeout=60s job/edc-proxy-migration-job
kubectl -n marketplace get pods
