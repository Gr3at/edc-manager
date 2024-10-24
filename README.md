
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

- Test all project folders: `go test ./... -json -coverprofile cover.out`
- Open coverage report in the browser: `go tool cover -html=cover.out`

- Bench testing: `go test ./... -bench=.`

- Race condition detector: `CGO_ENABLED=1 go test -race`

# Containerized Local Testing

`sh local-deployment.sh`


# K8S Deployment

To create the secret resource execute the following commands:
1. (if not present) `kubectl -n marketplace create secret generic edc-proxy-secrets --from-env-file=.env.prod --dry-run=client -o yaml`
2. (if not present) `kubectl -n marketplace create secret generic marketplace-registry-pull-secrets --from-file=.dockerconfigjson=.docker/config.json --type=kubernetes.io/dockerconfigjson --dry-run=client -o yaml`
3. (if not present) `sh scripts/deploy-postgres.sh`
4. (on every development update) `sh scripts/package.sh`
5. (on every development update) `sh scripts/publish.sh`
6. (to apply new database migrations - if any changes) `sh scripts/deploy-migration-job.sh`
7. (on every development update) `sh scripts/deploy-edc-proxy.sh`