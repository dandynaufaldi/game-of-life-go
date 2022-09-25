.PHONY: build

COVER_FILE="cover.out"

test:
	go test -coverprofile $(COVER_FILE) ./...
	go tool cover -func=$(COVER_FILE) | grep total
