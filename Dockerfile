FROM golang:1.25.1-trixie
LABEL tit="Web App Container"
ENV WORKDIR=/usr/src
WORKDIR /usr/src

RUN apt-get update

COPY . ${WORKDIR}

RUN go build -o main

ENV HOST=0.0.0.0 PORT=3000
EXPOSE 3000
ENTRYPOINT [ "./main" ]