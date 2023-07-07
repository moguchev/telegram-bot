LOCAL_BIN:=$(CURDIR)/bin


.PHONY: migrate
migrate:
	GOBIN=$(LOCAL_BIN) go install github.com/piiano/goose/v3/cmd/goose@v3.12.0 && \
    PATH=$(LOCAL_BIN):$(PATH) zsh ./migrate.sh