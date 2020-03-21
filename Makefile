CMD_ROOT=logger
PROJECT_NAME=logger
GIT_COMMIT=$$(git rev-parse --verify HEAD)
GIT_TAG=$$(git describe --tag $$(git rev-list --tags --max-count=1))
TIMESTAMP=$(shell date +'%Y%m%d%H%M%S')

-include ./makefile.properties

BIN_PATH=$(CMD_ROOT)_$$(go env GOOS)_$$(go env GOARCH)${BIN_EXT}

deps:
	go mod vendor -v
	go mod tidy -v
run:
	go run ./cmd/$(CMD_ROOT)
test:
	go test -v ./... -cover -coverprofile c.out
build:
	go build \
		-o ./bin/$(BIN_PATH) \
		./cmd/$(CMD_ROOT)
build_production:
	CGO_ENABLED=0 \
	go build -a -v \
		-ldflags "-X main.Commit=$(GIT_COMMIT) \
			-X main.Version=$(GIT_TAG) \
			-X main.Timestamp=$(TIMESTAMP) \
			-extldflags 'static' \
			-s -w" \
		-o ./bin/$(BIN_PATH) \
		./cmd/$(CMD_ROOT)
	sha256sum -b ./bin/$(BIN_PATH) \
		| cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
compress:
	ls -lah ./bin/$(BIN_PATH)
	upx -9 -v -o ./bin/.$(BIN_PATH) \
		./bin/$(BIN_PATH)
	upx -t ./bin/.$(BIN_PATH)
	rm -rf ./bin/$(BIN_PATH)
	mv ./bin/.$(BIN_PATH) \
		./bin/$(BIN_PATH)
	sha256sum -b ./bin/$(BIN_PATH) \
		| cut -f 1 -d ' ' > ./bin/$(BIN_PATH).sha256
	ls -lah ./bin/$(BIN_PATH)
see_ci:
	xdg-open https://gitlab.com/usvc/modules/go/logger/pipelines
.ssh:
	mkdir -p ./.ssh
	ssh-keygen -t rsa -b 8192 -f ./.ssh/id_rsa -q -N ""
	cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.base64
