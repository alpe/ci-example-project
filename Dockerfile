FROM scratch
COPY ci-example-project /

ENV port=8081
EXPOSE $port
ENTRYPOINT ["/ci-example-project"]
CMD ["-logtostderr=true"]
