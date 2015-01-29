.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  run         - run the program in dev mode"
	@echo "  test        - standard go test"
	@echo "  retest      - run tests, guard style"
	@echo "  build       - build the dist binary"
	@echo "  clean       - clean the dist build"
	@echo ""
	@echo "  tools       - go get's a bunch of tools for dev"
	@echo "  deps        - pull and setup dependencies"
	@echo "  update_deps - update deps lock file"

run:
	@(export CONFIG=$$PWD/godo.conf && cd ./cmd/godo-app-server && fresh -w=../..)

test:
	@go test ./... | grep -v "no test files" | sort -r

coverage:
	@go test -cover -v ./...

retest:
	@make test; reflex -r "^*\.go$$" -- make test

build: build_pkgs
	@mkdir -p ./bin
	@rm -f ./bin/*
	go build -o ./bin/godo-app-server github.com/pkieltyka/godo-app/cmd/godo-app-server

build_pkgs:
	go build ./... 

clean:
	@rm -rf ./bin

tools:
	go get github.com/robfig/glock
	go get github.com/cespare/reflex
	go get github.com/pkieltyka/fresh

deps:
	@glock sync -n github.com/pkieltyka/godo-app < Glockfile

update_deps:
	@glock save -n github.com/pkieltyka/godo-app > Glockfile
