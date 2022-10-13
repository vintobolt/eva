include .env

export PG_IP := ${PG_IP}
export PG_PORT := ${PG_PORT}
export PG_USER := $(PG_USER)
export PG_PASSWD := ${PG_PASSWD}
export PG_DBNAME := ${PG_DBNAME}
export SRV_IP := ${SRV_IP}
export SRV_PORT := ${SRV_PORT}
export SRV_LOGPATH := ${SRV_LOGPATH}
export SRV_GRACEFUL_TIMEOUT := ${SRV_GRACEFUL_TIMEOUT}
export SRV_LOGLEVEL := ${SRV_LOGLEVEL}

SWAG_PATH=/home/${USER}/Projects/eva/eva-server/app/cmd/app/main.go
SWAG_DOC=/home/${USER}/Projects/eva/eva-server/app/cmd/app/main.go

go_run:
	cd app && go run cmd/server/main.go
build:
	cd app && go build -o build/ cmd/server/main.go
run: build
	cd app && ./build/main
lint:
	cd app && golangci-lint run cmd/server/main.go
update-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
swagger:
	cd app && pwd && swag init -g cmd/server/main.go -o docs
