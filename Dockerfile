FROM ubuntu

ADD om /bin/

CMD [ "help" ]
ENTRYPOINT [ "/bin/om" ]