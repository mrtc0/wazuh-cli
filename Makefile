NAME := wazuh-cli

clean:
	-rm -rf build/

build:
	for os in darwin linux; do \
		for arch in amd64 386; do \
		  GOOS=$$os GOARCH=$$arch go build -o build/$$os-$$arch/$(NAME) wazuh-cli.go; \
		done;\
	done

install:
	dep ensure

update:
	dep ensure -update
