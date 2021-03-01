FROM golang:1.11.1
#Tao bien moi truong go
ENV GOPATH=/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

#Tiến hành setup
#Copy các file project vào container
COPY eureka /go/src/eureka-golang/eureka
COPY handle /go/src/eureka-golang/handle
COPY object /go/src/eureka-golang/object
COPY main.go /go/src/eureka-golang/main.go

#Thu muc lam viec
WORKDIR $GOPATH/src/eureka-golang/

RUN go get github.com/gorilla/mux

RUN go build main.go
# Command to run the executable
CMD ["./main"]