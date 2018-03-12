include ../build.mk

build: Dockerfile etc release
	@docker build --tag=$(REPONAME)/$(NAME)-rootfs:$(VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) \
		--build-arg=VERSION=$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-rootfs:$(VERSION)

run: build
	@docker run --interactive --tty --volume=$(CURDIR):/data \
		$(REPONAME)/$(NAME)-rootfs:$(VERSION) /bin/sh -l

release: release.go
	@go run release.go $(VERSION)

clean:
	@rm -f release
