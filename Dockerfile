FROM scratch
COPY demo /bin/demo
ENTRYPOINT ["/bin/demo"]