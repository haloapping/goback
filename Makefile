# Load variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# ===== GOOSE DATABASE MIGRATIONS ===== #
MIGRATIONS_DIR=./db/migration
DB_DRIVER=postgres
DB_URL=postgres://postgres:$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# create new migration
# usage: make create name=create_users_table
create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# apply all up migrations
# usage: make up
up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up

# rollback the last migration
# usage: make down
down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down

# redo last migration (down + up)
# usage: make redo
redo:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" redo

# reset all (down all, then up all)
# usage: make reset
reset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" reset

# print current DB version
# usage: make version
version:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" version

# show all migration status
# usage: make status
status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" status

# migrate up to a specific version
# usage: make up-to version=20240614120000
up-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up-to $(version)

# migrate down to a specific version
# usage: make down-to version=20240614120000
down-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down-to $(version)

# fix goose migration numbering if out of sync
# usage: make fix
fix:
	goose -dir $(MIGRATIONS_DIR) fix