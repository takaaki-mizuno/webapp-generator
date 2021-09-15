.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
.PHONY: test
test:
	go test -count 1 -p 1 -short -race -v ./...
.PHONY: lint
lint:
	go run golang.org/x/lint/golint -set_exit_status "./..."
.PHONY: build
build:
	go build