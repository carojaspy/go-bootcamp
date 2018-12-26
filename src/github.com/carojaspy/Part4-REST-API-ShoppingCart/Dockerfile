FROM golang:1.11.2
ADD . /go/src/github.com/carojaspy/Part4-REST-API-ShoppingCart
WORKDIR /go/src/github.com/carojaspy/Part4-REST-API-ShoppingCart
# RUN go get github.com/smartystreets/goconvey && go get -u github.com/go-sql-driver/mysql && go get github.com/astaxie/beego
EXPOSE 8000
CMD ["go", "run", "httpServer.go"]
