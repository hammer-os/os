ARG ALPINE_VERSION
ARG VERSION

FROM marraison/hammer-busybox:${VERSION} as busybox
FROM alpine:${ALPINE_VERSION} as overlay

# create rootfs namespace and install Alpine Linux binary packages
COPY --from=busybox / /pkg
RUN mkdir -p /rootfs/etc/apk && cp -r /etc/apk/* /rootfs/etc/apk/ \
	&& apk add --no-cache --initdb --allow-untrusted --root /rootfs \
		/pkg/x86_64/busybox-1.28.0-r0.apk \
	&& apk add --no-cache --root /rootfs \
		musl \
		busybox \
	&& true
#		busybox ca-certificates \
#		openssh-sftp-server \
#		openssh-server \
#		openssh-client \
#		btrfs-progs \
#		iptables \
#		device-mapper-libs \
#		libltdl \
#		libseccomp \
#		docker \
#	&& true

# remove directories/files with no use
#ARG NOT_WANT=" \
#	/rootfs/etc/apk /rootfs/lib/apk \
#	/rootfs/etc/conf.d \
#	/rootfs/etc/init.d \
#	/rootfs/etc/network \
#	/rootfs/etc/securetty \
#	/rootfs/usr/bin/findssl.sh \
#	/rootfs/usr/bin/iptables-xml \
#	/rootfs/usr/local \
#	/rootfs/var \
#"
#RUN rm -rf ${NOT_WANT}
RUN rm -rf /rootfs/etc/apk /rootfs/lib/apk /rootfs/usr

# create system and distribution relevant directories/files
ARG SYSTEM_MOUNTPOINTS="proc sys"
ARG MOUNTPOINTS="dev run tmp var"
#RUN mkdir -p /rootfs/home/adm && chmod 750 /rootfs/home/adm \
#	&& chown adm:adm /rootfs/home/adm \
RUN mkdir -p /rootfs/root && chmod 750 /rootfs/root \
	&& mkdir -p /rootfs/tmp && chmod 1777 /rootfs/tmp \
	&& for d in ${SYSTEM_MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 555 /rootfs/${d}; \
	done \
	&& for d in ${MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 755 /rootfs/${d}; \
	done

# install distribution configuration files
ADD etc /rootfs/etc

FROM scratch
COPY --from=overlay /rootfs /
