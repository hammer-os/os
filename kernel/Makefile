include ../build.mk

build: build-image

build-image:
	@docker build --tag=hammer-kernel:$(HAMMER_VERSION) \
		--build-arg=KERNEL_VERSION=$(KERNEL_VERSION) .

clean:
	@echo -n
