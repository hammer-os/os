include ../build.mk

build: Dockerfile init-$(ARCH)
	@docker build --tag=$(REPONAME)/$(NAME)-initrd:$(VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) \
		--build-arg=VERSION=$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-initrd:$(VERSION)

iso: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) iso \
		> ../hammer-$(ARCH).iso

kernel: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) kernel \
		> kernel-$(ARCH)

initrd: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) initrd \
		> initrd-$(ARCH).gz

init-x86_64:
	@make -C init rc.init

clean:
	@make -C init clean
	@rm -f hammer-*.iso kernel initrd rc.init
