
# EVA

api layer -> controller -> service -> repository -> db

# Migrations

`
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate
`

`migrate create -ext sql -dir ./migrations -seq init`
