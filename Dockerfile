FROM golang:alpine

LABEL maintener="VMS" version="1.0" description="ascii-art-web-dockerize"

LABEL taille="332MB"

RUN mkdir /build

WORKDIR /build

ADD . /build

RUN go build -o main .

RUN apk update && apk add bash && apk add tree

EXPOSE 8080

CMD [ "/build/main" ]