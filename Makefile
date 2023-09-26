TAG=`git describe --tags --abbrev=0`
DATE=`date +%FT%T%z`
HASH_COMMIT=$(shell git rev-parse HEAD)

ifeq (${HASH_COMMIT},HEAD)
	HASH_COMMIT=
endif

ifneq ($(shell git status --porcelain),)
	NOTE=build with uncommitted changes
endif

VERSION_PKG=github.com/MonstrHW/PoeTradeNotifier/internal/config

LDFLAGS=-ldflags "-w -s \
-X '${VERSION_PKG}.tag=${TAG}' \
-X '${VERSION_PKG}.date=${DATE}' \
-X '${VERSION_PKG}.hashCommit=${HASH_COMMIT}' \
-X '${VERSION_PKG}.note=${NOTE}'"

build_linux:
	go build ${LDFLAGS} -v -o build/PoeTradeNotifier cmd/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -v -o build/PoeTradeNotifier.exe cmd/main.go

test:
	go test -v ./...

.PHONY: build test
