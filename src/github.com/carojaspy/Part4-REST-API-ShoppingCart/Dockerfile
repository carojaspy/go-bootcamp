FROM golang:1.11.2
ADD . /go/src/github.com/carojaspy/Part4-REST-API-ShoppingCart
WORKDIR /go/src/github.com/carojaspy/Part4-REST-API-ShoppingCart
RUN  go get -u github.com/go-sql-driver/mysql && go get github.com/gorilla/mux
# Just to document the port to be exposed
EXPOSE 8000/tcp
# CMD ["go", "run", "httpServer.go"]
CMD ["./runserver.sh"]
