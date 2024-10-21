# make sure required secrets have been created
# kubectl -n marketplace create secret generic edc-proxy-secrets --from-env-file=.env.prod --dry-run=client -o yaml
# kubectl -n marketplace create secret generic marketplace-registry-pull-secrets --from-file=.dockerconfigjson=.docker/config.json --type=kubernetes.io/dockerconfigjson --dry-run=client -o yaml

echo "Automigrate database tables schema..."
export KUBECONFIG="/cygdrive/c/Users/dkaragkounis/Desktop/Omega-X/omega-x.yaml"
kubectl -n marketplace apply -f k8s/edc-proxy-migration-job.yaml
kubectl -n marketplace wait --for=condition="Complete" --timeout=60s job/edc-proxy-migration-job
kubectl -n marketplace get pods
