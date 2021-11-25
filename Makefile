run:
	go mod tidy && go mod download && GIN_MODE=debug CGO_ENABLED=0 CFG_POSTGRESCONNECTIONSTRING=postgres://postgres:mysecretpassword@localhost:5432/hw_database go run -tags migrate ./cmd