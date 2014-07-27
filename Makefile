SELFPKG := github.com/lmdtfy/lmdtfy
VERSION := 0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

.PHONY := test $(PKGS)

all: embed build

build:
	go build -o bin/lmdtfy -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/lmdtfy
	go build -o bin/lmdtfyd -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/lmdtfyd

build-dist: godep
	godep go build -o bin/lmdtfy -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/cmd/lmdtfy
	godep go build -o bin/lmdtfyd -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/cmd/lmdtfyd

bump-deps: deps vendor

deps:
	go get -u -t -v ./...

vendor: godep
	godep save ./...

test: $(PKGS)

$(PKGS): godep
	godep go test -v $@

# install:
# 	test -f /etc/default/drone || cp deb/drone/etc/default/drone /etc/default/drone
# 	cd bin && install -t /usr/local/bin drone
# 	cd bin && install -t /usr/local/bin droned
# 	mkdir -p /var/lib/drone


godep:
	go get github.com/tools/godep
