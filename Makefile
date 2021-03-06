run:
	go mod tidy && go mod download && GIN_MODE=debug CGO_ENABLED=0 CFG_POSTGRESCONNECTIONSTRING=postgres://postgres:mysecretpassword@host.docker.internal:5054/hw_database go run -tags migrate ./cmd

swagger:
	swag init -g internal/controller/http/apiv1/router.go