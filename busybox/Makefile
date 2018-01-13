include ../build.mk

build: build-busybox

build-busybox: Dockerfile
	@docker build --force-rm --tag=hammer-busybox:$(HAMMER_VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) .

menuconfig:
	@docker build --force-rm --tag=hammer-bb-conf:$(HAMMER_VERSION) \
		--build-arg=BUSYBOX_VERSION=$(BUSYBOX_VERSION) \
		--file Dockerfile.conf .
	@docker run --rm --interactive --tty --volume=$(CURDIR):/data \
		hammer-bb-conf:$(HAMMER_VERSION) sh

#	@docker run --rm hammer-busybox:$(HAMMER_VERSION) \
#		tar cC rootfs . | gzip -9 > \
#		hammer-busybox-$(HAMMER_VERSION).tar.gz

clean:
	@echo -n
