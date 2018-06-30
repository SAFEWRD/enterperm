# Copyright 2018 Safewrd Ventures OÜ
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0

VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)
VERSION_PACKAGE = $(REPOPATH)/pkg/version

GOOS ?= $(shell go env GOOS)
GOARCH = amd64
ORG := github.com/SAFEWRD
PROJECT := enterperm

REPOPATH ?= $(ORG)/$(PROJECT)

GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GO_LDFLAGS := '-extldflags "-static"
GO_LDFLAGS += -X $(VERSION_PACKAGE).version=$(VERSION)
GO_LDFLAGS += -w -s
GO_LDFLAGS += '

EXECUTOR_PACKAGE = $(REPOPATH)/cmd/enterperm

bin/enterperm: $(GO_FILES)
	GOARCH=$(GOARCH) GOOS=$(GOOS) CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ $(EXECUTOR_PACKAGE)

bin/enterperm-linux: $(GO_FILES)
	GOARCH=$(GOARCH) GOOS=linux CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ $(EXECUTOR_PACKAGE)

.PHONY: all
all: bin/enterperm bin/enterperm-linux

.PHONY: test
test: bin/enterperm
	go test -cover -v -timeout 60s `go list ./... | grep -v vendor`

.PHONY: clean
clean: 
	rm bin/*