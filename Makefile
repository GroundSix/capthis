BIN_PATH=bin/capthis

all: capthis

dependencies: vendor/
	git submodule init
	git submodule update
	git submodule foreach git pull origin master

capthis: dependencies main.go
	go build -o ${BIN_PATH} main.go