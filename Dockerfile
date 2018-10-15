FROM ubuntu:xenial

ADD om /bin/

CMD [ "help" ]
ENTRYPOINT [ "/bin/om" ]