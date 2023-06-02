DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl


# options: amd64 arm64
ARCH = arm64

# options: linux darwin windows
OS = darwin

# define the service names, these names are the same as the folder names inside the cmd directory
SERVICES := api user task experimental
service = $(word 1, $@)

# define the current node number
node = 0

# define the bin path
BIN = $(DIR)/bin

# Skywalking-go agent
BINARY = skywalking-go-agent
TOOLS_PATH = $(DIR)/tools
AGENT_SOURCE_PATH = $(DIR)/skywalking-go/tools/go-agent
AGENT_PATH = $(TOOLS_PATH)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH)
AGENT_CONFIG = $(DIR)/config/agent/agent.yaml

# go settings
GO = go
GO_BUILD = $(GO) build
GO_BUILD_FLAGS = -v
GO_BUILD_LDFLAGS = -X main.version=$(VERSION)

.PHONY: agent
agent:
	cd $(AGENT_SOURCE_PATH) && make deps
	cd $(AGENT_SOURCE_PATH) && \
	GOOS=$(OS) GOARCH=$(ARCH) $(GO_BUILD) $(GO_BUILD_FLAGS) -ldflags "$(GO_BUILD_LDFLAGS)" -o $(TOOLS_PATH)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH) ./cmd


.PHONY: proto
proto:
	@for file in $(IDL_PATH)/*.proto; do \
		protoc -I $(IDL_PATH) $$file --go-grpc_out=$(IDL_PATH)/pb --go_out=$(IDL_PATH)/pb; \
	done
	@for file in $(shell find $(IDL_PATH)/pb/* -type f); do \
		protoc-go-inject-tag -input=$$file; \
	done

.PHONY: $(SERVICES)
$(SERVICES):
	$(GO_BUILD) \
	-o $(BIN)/$(service) \
	-toolexec="$(AGENT_PATH) -config $(AGENT_CONFIG)" \
	$(CMD)/$(service)
	$(BIN)/$(service) -config $(CONFIG_PATH) -srvnum=$(node)

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down