API_PATH=api/weather
PROTO_OUT_DIR=pkg/weatherapi
PROTO_API_DIR=$(API_PATH)
ARGS=-fix

.PHONY: gen
gen: gen-proto generate

.PHONY: gen-proto
gen-proto:
	mkdir -p $(PROTO_OUT_DIR)
	protoc \
    	-I $(API_PATH)/v1 \
    	-I third_party \
    	--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
    	--go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(PROTO_OUT_DIR) --grpc-gateway_opt paths=source_relative \
      	--openapiv2_out=$(PROTO_OUT_DIR) --openapiv2_opt=allow_merge=true,merge_file_name=weather \
    ./$(PROTO_API_DIR)/v1/*.proto

run:
	go run cmd/main.go
evans:
	evans --port 9001 -r repl

test:
	go test ./... -cover -count=1

generate:
	go generate ./...

lint:
	go/lint proto/lint

go/lint:
	golangci-lint run  --config=.golangci.yml --timeout=30s ./...

proto/lint:
	protolint lint $(ARGS) $(PROTO_API_DIR)/*

run:
	go run cmd/weather/main.go

d-up:
	docker run --name=db -e POSTGRES_PASSWORD='qwerty' -p 5435:5432 -d --rm postgres

d-exec:
	docker exec -it grpc-app-db_v1.0 /bin/bash

m-up:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up
m-down:
	migrate -path ./scheme -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' down

apply:
	kubectl apply -f deployment/db-deployment.yml && \
    kubectl apply -f deployment/db-service.yml && \
    kubectl apply -f deployment/weather-deployment.yml && \
    kubectl apply -f deployment/weather-service.yml && \
    kubectl apply -f deployment/redis-deployment.yaml && \
    kubectl apply -f deployment/redis-service.yaml

# kubectl scale --replicas=0 deployment/<your-deployment>

# TODO: minikube ingress (open ports)
# TODO: deploy to kubernetes 1 replica postgres DB and 3 replicas of application
