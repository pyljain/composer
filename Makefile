export DB_TYPE=sqlite3
export DB_CONNECTION_STRING=composer.db

.PHONY: build-backend
build-backend:
	go build -o ./bin/composer

.PHONY: build-frontend
build-frontend:
	cd ui && \
	npm install && \
	npm run build

.PHONY: run-frontend
run-frontend:
	cd ui && \
	npm run dev

# .PHONY: run-backend
# run-backend: build-backend
#	./bin/composer

.PHONY: run-backend
run-backend:
	go run main.go

.PHONY: build-all
build-all: build-backend build-frontend
	
