#   Copyright 2020 The 2020 Docker, Inc.

#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at

#       http://www.apache.org/licenses/LICENSE-2.0

#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

PROTOS=$(shell find protos -name \*.proto)

EXTENSION:=
ifeq ($(GOOS),windows)
  EXTENSION:=.exe
endif

STATIC_FLAGS=CGO_ENABLED=0
LDFLAGS:="-s -w"
GO_BUILD=$(STATIC_FLAGS) go build -trimpath -ldflags=$(LDFLAGS)

BINARY?=bin/docker
BINARY_WITH_EXTENSION=$(BINARY)$(EXTENSION)

TAGS:=
ifdef BUILD_TAGS
  TAGS=-tags $(BUILD_TAGS)
endif

all: cli

protos:
	@protoc -I. --go_out=plugins=grpc,paths=source_relative:. ${PROTOS}

cli:
	GOOS=${GOOS} GOARCH=${GOARCH} $(GO_BUILD) $(TAGS) -o $(BINARY_WITH_EXTENSION) ./cli

cross:
	@GOOS=linux   GOARCH=amd64 $(GO_BUILD) $(TAGS) -o $(BINARY)-linux-amd64 ./cli
	@GOOS=darwin  GOARCH=amd64 $(GO_BUILD) $(TAGS) -o $(BINARY)-darwin-amd64 ./cli
	@GOOS=windows GOARCH=amd64 $(GO_BUILD) $(TAGS) -o $(BINARY)-windows-amd64.exe ./cli

test:
	@go test -count=1 $(TAGS) -cover $(shell go list ./... | grep -vE 'e2e')

lint:
	golangci-lint run --timeout 10m0s ./...

FORCE:

.PHONY: all protos cli cross test lint
