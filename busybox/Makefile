include ../build.mk

build: build-busybox

build-busybox: Dockerfile
	@docker build --tag=marraison/hammer-busybox:$(HAMMER_VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) .

menuconfig:
	@docker build --tag=hammer-busybox-conf:$(HAMMER_VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) \
		--file Dockerfile.conf .
	@docker run --rm --interactive --tty --volume=$(CURDIR):/data \
		hammer-busybox-conf:$(HAMMER_VERSION) /bin/sh -l

clean:
	@echo -n
