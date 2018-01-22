ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION} AS busybox-build

RUN apk update && apk upgrade && apk add --no-cache alpine-sdk

RUN adduser -G abuild -D abuild \
	&& mkdir -p busybox pkg var/cache/distfiles \
	&& chmod 775 busybox pkg var/cache/distfiles \
    && chown -R abuild:abuild busybox pkg var/cache/distfiles

COPY --chown=abuild:abuild . /busybox

WORKDIR /busybox
USER abuild
RUN abuild-keygen -i -a -n && abuild -r -P /pkg

FROM scratch
COPY --from=busybox-build /pkg /
