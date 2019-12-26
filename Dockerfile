FROM scratch
MAINTAINER Tony Grosinger <tony@grosinger.net>

COPY bin/simple-file-server /
EXPOSE 80
ENTRYPOINT ["/simple-file-server"]
