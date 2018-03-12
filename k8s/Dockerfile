ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION} AS kubernetes-build

RUN apk update && apk upgrade && apk add --no-cache alpine-sdk

RUN adduser -G abuild -D abuild \
	&& mkdir -p kubernetes pkg var/cache/distfiles \
	&& chmod 775 kubernetes pkg var/cache/distfiles \
    && chown -R abuild:abuild kubernetes pkg var/cache/distfiles

COPY --chown=abuild:abuild . /kubernetes

WORKDIR /kubernetes
USER abuild
RUN abuild-keygen -i -a -n && abuild -r -P /pkg

FROM scratch
COPY --from=kubernetes-build /pkg /
