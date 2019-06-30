FROM golang

WORKDIR /go/src/github.com
RUN git clone "https://github.com/ameena3/AnubhavWebsite.git"
WORKDIR /go/src/github.com/AnubhavWebsite
RUN go get -d -v ./...
RUN go build 

CMD ["main"]