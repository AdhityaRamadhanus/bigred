.PHONY: clean
# add test later 
# Flags #
GO_FLAGS = -v -race -o

# target #

default: clean build_bigred

quick_build: 
	go build -o bigred

build_bigred: 
	@echo "Build bigred"
	@go build $(GO_FLAGS) bigred
	@echo "...done"

