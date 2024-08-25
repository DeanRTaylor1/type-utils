SHELL := /bin/bash
GO_VERSION := 1.22.5
GO_BUILD := $(GO) build
Q := $(if $(filter 1,$(V)),,@)
M := $(shell printf "\033[36;1m>\033[0m")

.PHONY: all gen

gen:
	$(Q) cd grammar && ./generate.sh
	$(Q) cd ..
	@echo -e "$(M) Generated parser"
