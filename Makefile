include .env

export PG_IP := ${PG_IP}
export PG_PORT := ${PG_PORT}
export PG_USER := $(PG_USER)
export PG_PASSWD := ${PG_PASSWD}
export PG_DBNAME := ${PG_DBNAME}
export PG_TIMOUT := ${PG_TIMOUT}
export SRV_IP := ${SRV_IP}
export SRV_PORT := ${SRV_PORT}
export SRV_LOGPATH := ${SRV_LOGPATH}
export SRV_GRACEFUL_TIMEOUT := ${SRV_GRACEFUL_TIMEOUT}
export SRV_LOGLEVEL := ${SRV_LOGLEVEL}
export SRV_READ_TIMEOUT := ${SRV_READ_TIMEOUT}
export SRV_WRITE_TIMEOUT := ${SRV_WRITE_TIMEOUT}


go:
	go run cmd/api/main.go
build:
	go build -o bin/ cmd/api/main.go
run: build
	./bin/main
lint:
	golangci-lint run cmd/api/main.go
update-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
swagger:
	swag init -g cmd/api/main.go -o docs
migrate_up:
	migrate -path migrations -database "postgresql://$(PG_USER):$(PG_PASSWD)@$(PG_IP):$(PG_PORT)/$(PG_DBNAME)?sslmode=disable" -verbose up
migrate_down:
	migrate -path migrations -database "postgresql://$(PG_USER):$(PG_PASSWD)@$(PG_IP):$(PG_PORT)/$(PG_DBNAME)?sslmode=disable" -verbose down