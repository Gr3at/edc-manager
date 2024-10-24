set -e

IMAGE_TAG="${1:-v0.0.1}"

echo "Pushing image to registry..."
podman push --tls-verify=false registry.atosresearch.eu:18519/edc-proxy:${IMAGE_TAG}
echo "Image published to registry."
