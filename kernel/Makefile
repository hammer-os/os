include ../build.mk

build: build-image-$(ARCH)

build-image-x86_64: Dockerfile config-x86_64 keys.asc
	@docker build --tag=$(REPONAME)/$(NAME)-kernel:$(VERSION) \
		--tag=$(REPONAME)/$(NAME)-kernel:$(KERNEL_VERSION) \
		--build-arg=KERNEL_VERSION=$(KERNEL_VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-kernel:$(VERSION)

clean:
	@echo -n
