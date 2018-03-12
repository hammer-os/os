ARG ALPINE_VERSION
ARG VERSION
FROM marraison/hammer-rootfs:${VERSION} as rootfs
FROM marraison/hammer-kernel:${VERSION} as kernel
FROM alpine:${ALPINE_VERSION} as syslinux

RUN apk add --no-cache libarchive-tools cdrkit syslinux xz

COPY --from=rootfs / /rootfs
COPY --from=kernel / /kernel

# install syslinux bootloader
RUN mkdir -p /rootfs/boot/isolinux && chmod 755 /rootfs/boot/isolinux \
	&& cp /usr/share/syslinux/isolinux.bin /rootfs/boot/isolinux \
	&& cp /usr/share/syslinux/ldlinux.c32  /rootfs/boot/isolinux
ADD isolinux.cfg /rootfs/boot/isolinux/isolinux.cfg

# TODO: fix dead syslink at rootfs/lib/modules/4.14.13-hammer/build 
# install linux kernel and kernel modules
RUN mkdir -p /rootfs/boot && chmod 755 /rootfs/boot \
	&& cp /kernel/kernel /rootfs/boot/kernel \
	&& cp /kernel/System.map /rootfs/boot/System.map

# TODO: kernel headers and development files
#	&& tar xf /kernel/kernel-headers.tar -C /rootfs \
#	&& tar xf /kernel/kernel-dev.tar -C /rootfs \
RUN tar xf /kernel/kernel.tar -C /rootfs \
	&& chmod 755 /rootfs/lib \
	&& rm -f /rootfs/lib/modules/*/build

# install rc.(start|shutdown|...) handler (/sbin/init)
COPY rc.init /rootfs/etc/rc.init
COPY rc.dhcp /rootfs/etc/rc.dhcp
RUN chmod 755 /rootfs/etc/rc.init /rootfs/etc/rc.dhcp \
	&& chmod 444 /rootfs/etc/os-release

# install /etc/resolv.conf symlink and /etc/docker configuration
# directory (both must be writable), but /etc is read-only.
RUN ln -s /run/resolv.conf /rootfs/etc/resolv.conf

# create /etc/mtab symlink
RUN ln -s /proc/mounts /rootfs/etc/mtab

COPY mkinitrd.sh /bin/mkinitrd
RUN chmod 755 /bin/mkinitrd

WORKDIR /rootfs
ENTRYPOINT ["/bin/mkinitrd"]
CMD ["help"]
