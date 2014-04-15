BIN_PATH=bin/capthis
INSTALL_PATH=/usr/local/bin/capthis

all: capthis

dependencies: vendor/
	git submodule init
	git submodule update
	git submodule foreach git pull origin master

capthis: dependencies main.go
	go build -o ${BIN_PATH} main.go

install:
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed capthis!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled capthis"

test:
	@echo "Running tests:"
	go test capthis_test.go

clean:
	@echo "Deleting ${BIN_PATH}."
	rm -f ${BIN_PATH}