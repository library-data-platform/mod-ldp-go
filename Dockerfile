FROM iron/base

ADD ./mod-ldp /
ENTRYPOINT ["/mod-ldp"]
EXPOSE 8001
