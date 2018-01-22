include ../build.mk

build: Dockerfile init-$(ARCH)
	@echo PRETTY_NAME=\"HammerOS v${VERSION}\" > etc/os-release
	@echo NAME=\"HammerOS\" >> etc/os-release
	@echo ID=\"hammer\" >> etc/os-release
	@echo VERSION=\"v${VERSION}\" >> etc/os-release
	@docker build --tag=$(REPONAME)/$(NAME)-initrd:$(VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) \
		--build-arg=VERSION=$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-initrd:$(VERSION)

iso: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) iso \
		> hammer-$(ARCH).iso

kernel: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) kernel \
		> kernel-$(ARCH)

initrd: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) initrd \
		> initrd-$(ARCH).gz

init-x86_64:
	GOOS=linux GOARCH=amd64 go build -o rc.init rc.init.go

clean:
	@rm -f etc/os-release hammer-*.iso kernel initrd
	@go clean ./...
