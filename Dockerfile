FROM ubuntu:20.04

WORKDIR /app

RUN mkdir -p /app/.keystore
RUN mkdir -p /app/conf

COPY bin/delegator . 
COPY config.json conf
COPY .keystore .keystore

CMD ["/app/delegator"]
