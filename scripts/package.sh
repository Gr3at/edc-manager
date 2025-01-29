set -e

IMAGE_TAG="${1:-v0.0.1}"

echo "Building container image with tag ${IMAGE_TAG}..."
podman build -f build/package/Dockerfile -t localhost/edc-proxy:${IMAGE_TAG} .
podman tag localhost/edc-proxy:${IMAGE_TAG} registry.atosresearch.eu:18519/edc-proxy:${IMAGE_TAG}
echo "Container image built and tagging completed."
