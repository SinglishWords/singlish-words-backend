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

windows: build_dir config
	$(info [Windows_amd64] Building Windows distribution.)
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o .build/sw_windows_amd64.exe main.go
	$(info [Windows_amd64] Built Windows exe file to .build/sw_windows_amd64.exe)

linux: build_dir config
	$(info [Linux_amd64] Building Linux distribution.)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .build/sw_linux_amd64 main.go
	$(info [Linux_amd64] Built Linux exe file to .build/sw_linux_amd64)

run:
	go run main.go

swagger:
	$(info [Swagger] Init swagger docs.)
	~/go/bin/swag init --parseDependency --parseInternal

.PHONY: build
build: windows linux
	$(info [Build] Building done, see .build folder for distributions.)

.PHONY: all
all: clean swagger build
	$(info [All] Building done, see .build folder for distributions.)

#install_dirs:
#	$(info [Mkdirs] mkdir config dirs.)
#	$(info make /var/local/singlishwords/mysql)
#	@mkdir -p /var/local/singlishwords/mysql/
#	$(info make /var/local/singlishwords/www)
#	@mkdir -p /var/local/singlishwords/www
#	$(info make /var/local/singlishwords/log/web)
#	@mkdir -p /var/local/singlishwords/log/web
#	$(info make /var/local/singlishwords/log/nginx)
#	@mkdir -p /var/local/singlishwords/log/nginx
#	$(info make /etc/letsencrypt)
#	@mkdir -p /etc/letsencrypt
#	$(info made install dirs)

#install: install_dirs
#	@cp ./deployments/nginx/* /var/local/singlishwords/log/nginx/
#	@cp ./config.yaml /usr/local/etc/singlishwords/web/