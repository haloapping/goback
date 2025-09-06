# Load variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# === CONFIG ===
MIGRATIONS_DIR=./db/migration
DB_DRIVER=postgres
DB_URL=postgres://postgres:$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# === USAGE ===
# make create name=create_users_table
# make up
# make down
# make redo
# make reset
# make status
# make version
# make up-to version=20240614120000
# make down-to version=20240614120000
# make fix

# === COMMANDS ===

# Create new migration
create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Apply all up migrations
up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up

# Rollback the last migration
down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down

# Redo last migration (down + up)
redo:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" redo

# Reset all (down all, then up all)
reset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" reset

# Print current DB version
version:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" version

# Show all migration status
status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" status

# Migrate up to a specific version
up-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up-to $(version)

# Migrate down to a specific version
down-to:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down-to $(version)

# Fix goose migration numbering if out of sync
fix:
	goose -dir $(MIGRATIONS_DIR) fix