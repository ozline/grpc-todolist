DIR = $(shell pwd)/cmd

CONFIG_PATH = $(shell pwd)/config
IDL_PATH = $(shell pwd)/idl

SERVICES := api user task
service = $(word 1, $@)

BIN = $(shell pwd)/bin

.PHONY: proto
proto:
	@for file in $(IDL_PATH)/*.proto; do \
		protoc -I $(IDL_PATH) $$file --go-grpc_out=$(IDL_PATH)/pb --go_out=$(IDL_PATH)/pb; \
	done

.PHONY: $(SERVICES)
$(SERVICES):
	go build -o $(BIN)/$(service) $(DIR)/$(service)
	$(BIN)/$(service) -config $(CONFIG_PATH)

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down