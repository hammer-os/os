FROM marraison/hammer-busybox:0.0.1 as busybox-binary
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

# install busybox coreutils (provides /bin and /sbin)
RUN install -d -m 755 /rootfs/usr \
	&& cp -r sbin /rootfs/usr && chmod 755 /rootfs/usr/sbin \
	&& cp -r bin /rootfs/usr && chmod 755 /rootfs/usr/bin

# create /bin and /sbin symbolic links
WORKDIR /rootfs
RUN ln -s usr/bin bin && ln -s usr/sbin sbin

# create operating system rootfs
#ARG USRLAYOUT="usr usr/bin usr/sbin usr/share"
ARG ROOTLAYOUT="dev home lib mnt run var"
ARG USRLAYOUT=""
ARG BASELAYOUT="${ROOTLAYOUT} ${USRLAYOUT}"
ARG LINUXLAYOUT="proc sys"

RUN install -d -m 1777 /rootfs/tmp \
	&& for d in ${BASELAYOUT}; do \
		install -d -m 755 /rootfs/${d}; \
	done \
	&& for d in ${LINUXLAYOUT}; do \
		install -d -m 555 /rootfs/${d}; \
	done \
	&& install -d -m 700 /rootfs/root

# install hammer configuration scripts
ADD etc /rootfs/etc
RUN    find /rootfs/etc/ -type d -exec chmod 755 {} + \
	&& find /rootfs/etc/ -type f -exec chmod 644 {} + \
	&& chmod 755 /rootfs/etc

FROM scratch
COPY --from=busybox-binary /rootfs /
