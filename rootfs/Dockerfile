FROM marraison/hammer-busybox:0.0.1 as busybox-binary
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

# install busybox coreutils (provides /bin and /sbin)
RUN mkdir -p /rootfs/usr \
	&& cp -r sbin /rootfs/usr/sbin && chmod 755 /rootfs/usr/sbin \
	&& cp -r bin /rootfs/usr/bin && chmod 755 /rootfs/usr/bin

WORKDIR /rootfs
RUN ln -s usr/bin bin && ln -s usr/sbin sbin


# create operating system rootfs
#ARG USRLAYOUT="usr usr/bin usr/sbin usr/share"
ARG ""
ARG ROOTLAYOUT="dev home lib mnt run var"
ARG BASELAYOUT="${ROOTLAYOUT} ${USRLAYOUT}"
ARG LINUXLAYOUT="proc sys"

RUN mkdir /rootfs/tmp && chmod 1777 /rootfs/tmp \
	&& for d in ${BASELAYOUT} ${LINUXLAYOUT}; do \
		mkdir -p /rootfs/${d}; \
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
