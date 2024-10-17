echo "Creating podman network..."
podman network create edc-network
echo "Starting db container..."
podman run --name edc-db --network edc-network -p 5432:5432/tcp -e POSTGRES_PASSWORD=test -e POSTGRES_USER=test -e POSTGRES_DB=edc -d postgres:15.1-alpine3.16
echo "Starting edc-proxy container..."
podman run --name edc-proxy --network edc-network -p 8080:8080/tcp --env-file .env.prod -d edc-proxy:v0.0.1
echo "Local deployment ready."
