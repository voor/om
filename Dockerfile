FROM gcr.io/distroless/base

ADD om /bin/

CMD [ "help" ]
ENTRYPOINT [ "/bin/om" ]