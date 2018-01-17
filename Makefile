include build.mk

build:
	@make -C busybox build
	@make -C overlay build
	@make -C rootfs build

clean:
	@make -C busybox clean
	@make -C overlay clean
	@make -C rootfs clean
	@make -C kernel clean
	@make -C initrd clean
	
