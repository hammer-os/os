FROM marraison/hammer-busybox:0.0.1 as busybox-binary

# install busybox coreutils (provides /bin and /sbin)
RUN mkdir /rootfs \
	&& cp -r sbin /rootfs/sbin && chmod 755 /rootfs/sbin \
	&& cp -r bin /rootfs/bin && chmod 755 /rootfs/bin

# create operating system rootfs
ARG BASELAYOUT="dev home lib mnt run var"
ARG LINUXLAYOUT="proc sys"

RUN mkdir /rootfs/tmp && chmod 1777 /rootfs/tmp \
	&& for d in ${BASELAYOUT} ${LINUXLAYOUT}; do \
		mkdir /rootfs/${d}; \
	done \
	&& for d in ${BASELAYOUT}; do \
		chmod 755 /rootfs/${d}; \
	done \
	&& for d in ${LINUXLAYOUT}; do \
		chmod 555 /rootfs/${d}; \
	done \
	&& mkdir /rootfs/root && chmod 700 /rootfs/root

# install hammer configuration scripts
ADD etc /rootfs/etc
RUN    find /rootfs/etc/ -type d -exec chmod 755 {} + \
	&& find /rootfs/etc/ -type f -exec chmod 644 {} + \
	&& chmod 755 /rootfs/etc

FROM scratch
COPY --from=busybox-binary /rootfs /
