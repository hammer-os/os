include build.mk

build:
	@make -C busybox build
	@make -C rootfs build
	@make -C initrd build

clean:
	@make -C busybox clean
	@make -C rootfs clean
	@make -C kernel clean
	@make -C initrd clean
	
