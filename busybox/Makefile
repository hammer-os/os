include ../build.mk

build: Dockerfile config keys.asc
	@docker build --tag=$(REPONAME)/$(NAME)-busybox:$(VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-busybox:$(VERSION)

menuconfig: Dockerfile.conf config keys.asc
	@docker build --tag=$(NAME)-busybox-conf:$(VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) \
		--file Dockerfile.conf .
	@docker run --rm --interactive --tty \
			--volume=$(CURDIR):/data \
		$(NAME)-busybox-conf:$(VERSION) sh

clean:
	@echo -n
