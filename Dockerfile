FROM golang
WORKDIR /go/src/github.com/AnubhavWebsite
COPY . .
RUN go get -d -v ./...
ENV GOBIN=/go/src/github.com/AnubhavWebsite
ENV PATH=${PATH}:/go/src/github.com/AnubhavWebsite
RUN go install main.go
CMD ["main"]