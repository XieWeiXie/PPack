COMMIT  := $(shell git log -n 1 --pretty=format:"%h")
BRANCH  := $(shell git name-rev --name-only HEAD)
NOW     := $(shell date +%F_%T)

VERSION ?= v1.0.0-$(COMMIT)
VARS    := -X main.URL

ifeq ($(strip $(DEBUG)), true)
        BUILD_ARGS = -v -gcflags="all=-N -l" -ldflags '$(VARS)'
else
        BUILD_ARGS = -v -trimpath \
            -gcflags="all=-trimpath=$(PWD)" \
            -asmflags="all=-trimpath=$(PWD)" \
            -ldflags '$(VARS)'
endif

GOBUILD=go build $(BUILD_ARGS)

export GO111MODULE=on


