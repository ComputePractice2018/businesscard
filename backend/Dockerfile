FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/ComputePractice2018/businesscard/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/businesscard github.com/ComputePractice2018/businesscard/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/businesscard /opt/businesscard

EXPOSE 8080

ENV CONNECTION "businesscard:SuperSecretPassword@tcp(db:3306)/businesscard"
ENTRYPOINT /opt/businesscard -connection ${CONNECTION}