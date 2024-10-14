
# Get started

- Run the service: `go mod tidy` & `go run main.go`

# Production

- Build: `go build .`
- Run executable: `./edc-proxy`

# Development

- Make sure all dependencies are installed: `go mod tidy`

# Testing

- Test all project folders: `go test ./... -v`

- Bench testing: `go test -bench=.`

- Race condition detector: `CGO_ENABLED=1 go test -race`