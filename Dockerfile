FROM golang:1.11.1
#Tao bien moi truong go
ENV GOPATH=/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

#Tiến hành setup
#Copy các file project vào container
COPY eureka /go/src/demo-eureka/eureka
COPY handle /go/src/demo-eureka/handle
COPY model  /go/src/demo-eureka/model
COPY service /go/src/demo-eureka/service
COPY templates /go/src/demo-eureka/templates
COPY util /go/src/demo-eureka/util
COPY main.go /go/src/demo-eureka/main.go

#Thu muc lam viec
WORKDIR $GOPATH/src/demo-eureka/

RUN go get github.com/gorilla/mux
RUN go get github.com/satori/go.uuid

RUN go build main.go
# Command to run the executable
CMD ["./main"]