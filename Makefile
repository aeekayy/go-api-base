BINARY_NAME = go-api-base
OSX_BUILD_FLAGS := -s
STATIC_BUILD_FLAGS := -linkmode external -extldflags -static -w
STATIC_OSX_BUILD_FLAGS := -linkmode external -extldflags -w -s
ARCH := $(shell uname -m)

ifeq ($(GIT_COMMIT),)
        GIT_COMMIT=$(shell git rev-parse HEAD)
endif

BUILD_INFO_FLAGS := -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S') -X main.BuildCommitHash=$(GIT_REVISION)
BUILD_INFO_FLAGS_GIT := -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S') -X main.BuildCommitHash=$(GIT_COMMIT)

BINARY_NAME_GIT = $(BINARY_NAME)-$(GIT_COMMIT)

define BUILD_MANIFEST
git_repo: ${GIT_REPO}
git_branch: ${GIT_BRANCH}
git_commit: ${GIT_COMMIT}
git_owner: VeritoneEnergy
author: ${GIT_AUTHOR}
jenkins_url: ${RUN_DISPLAY_URL}
jenkins_build: ${BUILD_NUMBER}
version: ${VERSION_NUM}
endef
export BUILD_MANIFEST

gen-manifest: ## Generate a Build manifest
	@echo "$${BUILD_MANIFEST}" > build-manifest.yml

.PHONY: binary-build
binary-build:
	# Mac
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -ldflags "$(BUILD_INFO_FLAGS) $(OSX_BUILD_FLAGS)" -a -o bin/$(BINARY_NAME_GIT)-darwin-amd64
	# AMD64
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_INFO_FLAGS) $(OSX_BUILD_FLAGS)" -a -o bin/$(BINARY_NAME_GIT)-linux-amd64

.PHONY: build-docker-binary
build-linux-binary: gen-manifest
	[ ! -d bin ] && mkdir bin || echo "bin directory already available"
	go build -ldflags "$(BUILD_INFO_FLAGS) $(STATIC_BUILD_FLAGS)" -a -o bin/$(BINARY_NAME) .

.PHONY: build
build:
	docker build -t go-api-base .
