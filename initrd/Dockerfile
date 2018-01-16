FROM marraison/hammer-rootfs:0.0.1 as rootfs
FROM marraison/hammer-kernel:0.0.1 as kernel
FROM alpine:3.7 as syslinux
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

RUN apk add --no-cache libarchive-tools cdrkit syslinux xz

COPY --from=rootfs / /rootfs
COPY --from=kernel / /kernel

# install syslinux bootloader
RUN mkdir -p /rootfs/boot/isolinux && chmod 755 /rootfs/boot/isolinux \
	&& cp /usr/share/syslinux/isolinux.bin /rootfs/boot/isolinux/isolinux.bin \
	&& cp /usr/share/syslinux/ldlinux.c32  /rootfs/boot/isolinux/ldlinux.c32
ADD isolinux.cfg /rootfs/boot/isolinux/isolinux.cfg

# TODO: fix dead syslink at rootfs/lib/modules/4.14.13-hammer/build 
# install kernel and kernel modules
RUN mkdir -p /rootfs/boot && chmod 755 /rootfs/boot \
	&& cp /kernel/kernel /rootfs/boot/kernel \
	&& cp /kernel/System.map /rootfs/boot/System.map

# TODO: install kernel-headers and kernel-dev??? if not remove build
# link.
#	&& tar xf /kernel/kernel-headers.tar -C /rootfs \
#	&& tar xf /kernel/kernel-dev.tar -C /rootfs \
RUN tar xf /kernel/kernel.tar -C /rootfs \
	&& chmod 755 /rootfs/lib \
	&& rm -f /rootfs/lib/modules/*/build


# install system specific configuration
COPY etc /rootfs/etc
RUN find /rootfs/etc/ -type d -exec chmod 755 {} + \
	&& find /rootfs/etc/ -type f -exec chmod 644 {} + \
	&& chmod 755 /rootfs/etc

# install rc.(start|shutdown|...) handler (/sbin/init)
COPY rc.init /rootfs/etc/rc.init
COPY rc.dhcp /rootfs/etc/rc.dhcp
RUN chmod 755 /rootfs/etc/rc.init /rootfs/etc/rc.dhcp


# install /etc/resolv.conf symlink (must be writable, but /etc is
# read-only.
RUN ln -s /run/resolv.conf /rootfs/etc/resolv.conf


COPY mkinitrd.sh /bin/mkinitrd
RUN chmod 755 /bin/mkinitrd

WORKDIR /rootfs
ENTRYPOINT ["/bin/mkinitrd"]
CMD ["help"]
