.PHONY: app build binary default clear

COMMIT  := $(shell git log -n 1 --pretty=format:"%h")
BRANCH  := $(shell git name-rev --name-only HEAD)
NOW     := $(shell date +%F_%T)

VERSION ?= v1.0.0-$(COMMIT)
URL ?=https://www.douyin.com/
VARS    := -X main.URL=$(URL)

APP_NAME ?= app
ICON_NAME ?= app.icns
APP_VARS := -X main.APPName=$(APP_NAME) -X main.ICONName=$(ICON_NAME)

BUILD_ARGS = -v -gcflags="all=-N -l" -ldflags '$(VARS)'
APP_RUN_ARGS = -v -gcflags="all=-N -l" -ldflags '$(APP_VARS)'

export GO111MODULE=on

GO_RUN_BUILD := go build $(BUILD_ARGS) -o $(APP_NAME) ./binary/main.go
PACK_RUN_BINARY := go run $(APP_RUN_ARGS) main.go
CLEAR_RUN := rm $(APP_NAME)

default: app

app: build
	$(PACK_RUN_BINARY)
	$(CLEAR_RUN)

build:
	$(GO_RUN_BUILD)

clear:
	$(CLEAR_RUN)
