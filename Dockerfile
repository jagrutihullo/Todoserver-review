FROM golang

WORKDIR /go/src/todoserver
COPY * /go/src/todoserver/
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/mysql
RUN go get github.com/go-sql-driver/mysql
RUN go build
RUN go install
EXPOSE 9000
ENTRYPOINT todoserver
