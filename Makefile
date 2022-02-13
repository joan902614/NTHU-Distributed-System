PATH := $(CURDIR)/bin:$(PATH)

MODULES := video comment

DOCKER_COMPOSE := docker compose

####################################################################################################
### Automatically include components' extensions and ad-hoc rules (makefile.mk)
###
-include */makefile.mk

####################################################################################################
### Rule for the `generate` command
###

define make-dc-generate-rules

.PHONY: dc.$1.generate

# to generate individual module, override the command defined in the docker-compose.yml file
dc.$1.generate:
	$(DOCKER_COMPOSE) run --rm proto make $1.generate

endef
$(foreach module,$(MODULES),$(eval $(call make-dc-generate-rules,$(module))))

.PHONY: dc.generate
dc.generate:
	$(DOCKER_COMPOSE) run --rm generate

define make-generate-rules

$1.generate: bin/protoc-gen-go bin/protoc-gen-go-grpc bin/mockgen
	protoc \
		-I . \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		./modules/$1/pb/*.proto

	go generate ./modules/$1/...

endef
$(foreach module,$(MODULES),$(eval $(call make-generate-rules,$(module))))

generate: $(addsuffix .generate,$(MODULES))

bin/protoc-gen-go: go.mod
	go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-go-grpc: go.mod
	go build -o $@ google.golang.org/grpc/cmd/protoc-gen-go-grpc

bin/mockgen: go.mod
	go build -o $@ github.com/golang/mock/mockgen

####################################################################################################
### Rule for the `test` command
###

define make-dc-test-rules

.PHONY: dc.$1.test

# to test individual module, override the command defined in the docker-compose.yml file
dc.$1.test:
	$(DOCKER_COMPOSE) run --rm test make $1.test

endef
$(foreach module,$(MODULES),$(eval $(call make-dc-test-rules,$(module))))

.PHONY: dc.test
dc.test:
	$(DOCKER_COMPOSE) run --rm test

define make-test-rules

$1.test:
	go test -v -race ./modules/$1/...
	
endef
$(foreach module,$(MODULES),$(eval $(call make-test-rules,$(module))))

test: $(addsuffix .test,$(MODULES))
