include build.mk

build:
	@make -C busybox build
	@make -C kernel build

clean:
	@make -C busybox clean
	@make -C kernel clean
