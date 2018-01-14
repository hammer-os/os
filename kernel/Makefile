include ../build.mk

build: build-image-amd64

build-image-amd64: Dockerfile config-amd64 keys.asc
	@docker build --tag=$(REPONAME)/$(NAME)-kernel:$(VERSION) \
		--build-arg=KERNEL_VERSION=$(KERNEL_VERSION) .

clean:
	@echo -n
