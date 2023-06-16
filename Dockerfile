FROM golang:1.19-alpine

LABEL Maintener(captain)="Seynabou NIANG <https://learn.zone01dakar.sn/git/sniang>"
LABEL Maintener(backend)="Vincent Félix NDOUR <https://learn.zone01dakar.sn/git/vindour>"
LABEL Maintener(frontend)="Masseck THIAW <https://learn.zone01dakar.sn/git/mthiaw>"
LABEL Name="Ascii-Art-Web-Dockerize"
LABEL documentation="This Web app allows you to generate the graphic representation of all the printable ASCII-This is the dockerize version of it"
LABEL README="<https://learn.zone01dakar.sn/git/sniang/ascii-art-web-dockerize/src/branch/master/README.md>"
LABEL version="1.0.0"
LABEL License="VMS team"

RUN mkdir /build

WORKDIR /build

ADD . /build

RUN go build -o main .

RUN apk update && apk add bash && apk add tree

EXPOSE 8080

CMD [ "/build/main" ]

