#
# Tools and binaries
#
DOCKERCOMPOSECMD	= docker-compose
GOCMD			= go
GOTEST			=$(GOCMD) test
PROTOC			= protoc

#
# Directories and packages
#
TEST_PKGS := $(shell go list ./... | grep -v fixtures)

.PHONY: docker-compose
docker-compose:
	$(DOCKERCOMPOSECMD) up -d

.PHONY: generate-report
generate-report:
	cd ./example; yarn start

.PHONY: stats
stats:
	$(GOCMD) run main.go stats -d ./example

.PHONY: elastic
elastic:
	$(GOCMD) run main.go elastic -d ./example


.PHONY: test
test:
	$(GOTEST) $(TEST_PKGS)

.PHONY: testv
testv:
	$(GOTEST) -v $(TEST_PKGS)

