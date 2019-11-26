FROM alpine:3.10.3
COPY demo /bin/demo
ENTRYPOINT ["/bin/demo"]