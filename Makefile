.PHONY: app build binary default clear

COMMIT  := $(shell git log -n 1 --pretty=format:"%h")
BRANCH  := $(shell git name-rev --name-only HEAD)
NOW     := $(shell date +%F_%T)

VERSION ?= v1.0.0-$(COMMIT)
URL ?=https://www.douyin.com/
VARS    := -X main.URL=$(URL)

APPNAME ?= app
ICONNAME ?= app.icns
APPVARS := -X main.APPName=$(APPNAME) -X main.ICONName=$(ICONNAME)

BUILD_ARGS = -v -gcflags="all=-N -l" -ldflags '$(VARS)'
APP_RUN_ARGS = -v -gcflags="all=-N -l" -ldflags '$(APPVARS)'

export GO111MODULE=on

GO_RUN_BUILD := go build $(BUILD_ARGS) -o $(APPNAME) ./binary/main.go
PPAACk_RUN_BINARY := go run $(APP_RUN_ARGS) main.go
CLEAR_RUN := rm $(APPNAME)

default: app

app: build
	$(PPAACk_RUN_BINARY)
	$(CLEAR_RUN)

build:
	$(GO_RUN_BUILD)

clear:
	$(CLEAR_RUN)
