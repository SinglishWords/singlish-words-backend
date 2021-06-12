.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

build_dir:
	$(info [build_dir] Making temp building dir.)
	@mkdir -p .build
	$(info [build_dir] Made temp building dir as .build.)

.PHONY: clean
clean:
	$(info [Clean] Cleaning build folder.)
	@rm -rf .build
	$(info [Clean] Cleaned build folder.)

config: build_dir
	@cp config.yaml .build/

windows: build_dir config swagger
	$(info [Windows_amd64] Building Windows distribution.)
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o .build/main_windows_amd64.exe main.go
	$(info [Windows_amd64] Built Windows exe file to .build/main_windows_amd64.exe)

linux: build_dir config swagger
	$(info [Linux_amd64] Building Linux distribution.)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .build/main_linux_amd64 main.go
	$(info [Linux_amd64] Built Linux exe file to .build/main_linux_amd64)

run: swagger
	go run main.go

swagger:
	$(info [Swagger] Init swagger docs.)
	~/go/bin/swag init --parseDependency --parseInternal

.PHONY: build
build: clean swagger windows linux
	$(info [All] Building done, see .build folder for distributions.)