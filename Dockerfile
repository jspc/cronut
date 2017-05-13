FROM alpine
MAINTAINER jspc <james@zero-internet.org.uk>

ADD cronut-linux /cronut

ENTRYPOINT ["/cronut"]
