
# Get started

- Run the service: `go mod tidy` & `go run main.go`

# Production

- Migrate DB: `go run migrate/migrate.go`
- Build: `go build .`
- Run executable: `./edc-proxy`

# Development

- Make sure all dependencies are installed: `go mod tidy`
- Start postgres db: `podman run --name edc-db -p 5432:5432/tcp -e POSTGRES_PASSWORD=test -e POSTGRES_USER=test -e POSTGRES_DB=edc -d postgres:15.1-alpine3.16`

# Testing

- Test all project folders: `go test ./... -v`

- Bench testing: `go test -bench=.`

- Race condition detector: `CGO_ENABLED=1 go test -race`

# Containerized Local Testing

`sh local-deployment.sh`


# K8S Deployment

To create the secret resource execute the following commands:
1. `kubectl create secret generic app-secrets --from-env-file=.env.prod --dry-run=client -o yaml`
2. `kubectl create secret generic regcred --from-file=.dockerconfigjson=config.json --type=kubernetes.io/dockerconfigjson --dry-run=client -o yaml`
