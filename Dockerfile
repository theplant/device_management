FROM busybox
ADD views /gopkg/src/github.com/qor/qor/admin/views
ENV GOPATH /gopkg
ADD devicem /bin/
CMD /bin/devicem
