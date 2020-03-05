.DEFAULT_GOAL = help

BIN ?= $(shell which ./gs-server)

help:
	@cat $(MAKEFILE_LIST) | \
		grep -E '^# ~~~ .*? [~]+$$|^[a-zA-Z0-9_-]+:.*?## .*$$' | \
		awk '{if ( $$1=="#" ) {\
			match($$0, /^# ~~~ (.+?) [~]+$$/, a);\
			{print "\n", a[1], "\n"}\
		} else { \
			match($$0, /^([a-zA-Z-]+):.*?## (.*)$$/, a); \
			{printf "  - \033[32m%-20s\033[0m %s\n",   a[1], a[2]} \
		};}'
	@echo  ""

# ~~~ Go Lang Size Of ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
test:
	echo $(BIN)

build: ## Build Server
	go build -ldflags "-w -s" -o "gs-server"

start: ## Start Server
	$(BIN) start

stop: ## Stop Server
	$(BIN) stop

status: ## Show Server status
	$(BIN) status

restart: ## Restart service
	$(BIN) restart
