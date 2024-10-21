set -e
echo "Building container image..."
podman build -t localhost/edc-proxy:v0.0.1 .
podman tag localhost/edc-proxy:v0.0.1 registry.atosresearch.eu:18519/edc-proxy:v0.0.1
echo "Container image built and tagging completed."
