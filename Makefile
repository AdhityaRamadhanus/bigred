.PHONY: clean
# add test later 
# Flags #
GO_FLAGS = -race -o

# target #

default: clean build_bigred

quick_build: 
	go build -o bigred

build_bigred: 
	go build $(GO_FLAGS) bigred

# the executable will be named bigred, always 
clean:
	-rm bigred
