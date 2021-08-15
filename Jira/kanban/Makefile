REGISTRY_HOST = registry.gitlab.com
IMAGE = leanlabs/kanban
TAG   = 1.7.2
CWD   = /go/src/gitlab.com/leanlabsio/kanban
PACKAGES := `docker run --rm \
				-v $(CURDIR):$(CWD) \
				-w $(CWD) \
				golang:1.8-alpine go list ./... | grep -v '/vendor/'`

all: clean

test: rel/kanban_x86_64_linux
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		golang:1.8-alpine go test $(PACKAGES)

node_modules/: package.json
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-v $$HOME/node_cache:/cache \
		-w $(CWD) \
		-e HOME=/cache \
		node:6.3.0-wheezy npm run install

build: node_modules/
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-v $$HOME/node_cache:/cache \
		-w $(CWD) \
		-e HOME=/cache \
		node:6.3.0-wheezy npm run build

templates/templates.go: $(find $(CURDIR)/templates -name "*.html" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=templates -o templates/templates.go \
		templates/...

web/web.go: $(find $(CURDIR)/web/ -name "*" ! -name "web.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=web -o web/web.go \
		web/assets/... web/images/... web/template/...

rel/kanban_x86_64_linux: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.8.0 build -ldflags "-X main.AppVer=$(TAG) -s" -v -o $@

rel/kanban_x86_64_darwin: clean build templates/templates.go web/web.go $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		-e GO15VENDOREXPERIMENT=1 \
		-e CGO_ENABLED=0 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.8.0 build -ldflags "-X main.AppVer=$(TAG) -s" -v -o $@

login:
	@docker login -u $(REGISTRY_USERNAME) -p $(REGISTRY_PASSWORD) $(REGISTRY_HOST)

release: login rel/kanban_x86_64_linux
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

clean:
	@-docker run \
	    -v $(CURDIR):$(CWD) \
	    -w $(CWD) \
		alpine:3.4 rm -rf web/

	@-docker run \
		-v $(CURDIR):$(CWD) \
    	-w $(CWD) \
    	alpine:3.4 rm -f templates/templates.go

	@-docker run \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		alpine:3.4 rm -rf rel node_modules

# Development targets
dev_redis:
	@docker inspect -f {{.State.Running}} kb_dev_redis || \
		docker run -d -p 6379:6379 --name kb_dev_redis leanlabs/redis

watch: build
	@docker inspect -f {{.State.Running}} kb_dev_watcher || \
		docker run -d \
			--name kb_dev_watcher \
			-v $(CURDIR):$(CWD) \
			-v $$HOME/node_cache:/cache \
			-w $(CWD) \
			-e HOME=/cache \
			node:6.3.0-wheezy npm run watch

tmp/go/pkg/:
	@docker run --rm \
		-v $(CURDIR)/tmp/go/pkg:/go/pkg \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.8.0 install -v

dev : DEBUG=-debug

dev: watch templates/templates.go web/web.go dev_redis tmp/go/pkg/
	-docker rm -f kb_dev
	@docker run -d --link kb_dev_redis:redis --name kb_dev \
		-p 9000:80 \
		-v $(CURDIR):$(CWD) \
		-v $(CURDIR)/tmp/go/pkg:/go/pkg \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		-e KANBAN_GITLAB_URL=$(KANBAN_GITLAB_URL) \
		-e KANBAN_ENABLE_SIGNUP=$(KANBAN_ENABLE_SIGNUP) \
		-e KANBAN_SERVER_HOSTNAME=$(KANBAN_SERVER_HOSTNAME) \
		-e KANBAN_GITLAB_CLIENT=$(KANBAN_GITLAB_CLIENT) \
		-e KANBAN_GITLAB_SECRET=$(KANBAN_GITLAB_SECRET) \
		-e KANBAN_REDIS_ADDR=redis:6379 \
		-e KANBAN_AUTO_COMMENTS=false \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.8.0 run -v main.go server

.PHONY: help test build release
