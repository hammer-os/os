include ../build.mk

build: Dockerfile APKBUILD busybox.config busybox.post-install
	@docker build --tag=$(REPONAME)/$(NAME)-busybox:$(VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-busybox:$(VERSION)

clean:
	@echo -n
