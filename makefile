GO111MODULE=on

PROJECTNAME = $(shell basename "$(PWD)")
GOBASE = $(shell pwd)
MAINFILE = cmd/main.go

.PHONY: build-vue
build-vue:
	@echo "  > start to build vue files to dist ......"
	npm install
	npm run build
	@echo "  < vue build done !"
	go generate $(MAINFILE)

.PHONY: build-go
build-go:
	go generate $(MAINFILE)
	go build $(RACE) -o bin/$(PROJECTNAME) $(MAINFILE)
	@echo "  < go build done !"

.PHONY: run
run:
	go run $(MAINFILE) --config=./config/config.ini
