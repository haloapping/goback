# Load variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# ===== DATABASE SEEDER ===== #
dbseed:
	go run db/seeder/seed.go --nuser=$(nuser) --ntask=$(ntask)

# ===== GOOSE DATABASE MIGRATIONS ===== #
MIGRATIONS_DIR=./db/migration
DB_DRIVER=postgres
DB_URL=postgres://postgres:$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# create new migration
# usage: make dbcreate name=create_users_table
dbcreate:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# apply all up migrations
# usage: make dbup
dbup:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up

# rollback the last migration
# usage: make dbdown
dbdown:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down

# redo last migration (down + up)
# usage: make dbredo
dbredo:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" redo

# reset all (down all, then up all)
# usage: make dbreset
dbreset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up

# print current DB version
# usage: make dbversion
dbversion:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" version

# show all migration status
# usage: make dbstatus
dbstatus:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" status

# migrate up to a specific version
# usage: make dbup-to version=20240614120000
dbup-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up-to $(version)

# migrate down to a specific version
# usage: make dbdown-to version=20240614120000
dbdown-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down-to $(version)

# fix goose migration numbering if out of sync
# usage: make dbfix
dbfix:
	goose -dir $(MIGRATIONS_DIR) fix