FROM alpine:edge
MAINTAINER jspc <james@zero-internet.org.uk>

RUN apk add --update ca-certificates \
                     libgit2 \
                     libgit2-dev

ADD cronut /cronut

ENTRYPOINT ["/cronut"]
