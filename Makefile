APP_NAME:=admlte
## set the following env before run make file
#DB_HOST=
#DB_USER=
#DB_PWD=
#DB_DATABASE=
#DB_PORT=


.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go


.ONESHELL:
.PHONY: model-gen 
model-gen: $(eval SHELL:=/bin/bash)
	cd internal/store
	rm -f sql/queries.sql
	bash -c 'cat sql/queries/*.sql > sql/queries.sql'
	sqlc generate

.PHONY: pgosql
pgosql:
	go build -o ./bin/pgosql ./cmd/pgosql


.PHONY: db-init
db-init: 
	go run ./cmd/pgosql/main.go -user $(DB_USER) -password $(DB_PWD) -host $(DB_HOST) -port $(DB_PORT) -db $(DB_DATABASE) -file internal/store/sql/schema.sql 
	
.PHONY: db-fake
db-fake: 
	go run ./cmd/faker/main.go -user $(DB_USER) -password $(DB_PWD) -host $(DB_HOST) -port $(DB_PORT) -db $(DB_DATABASE) 