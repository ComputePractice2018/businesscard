FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/businesscard/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN go get -u github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/businesscard github.com/ComputePractice2018/businesscard/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/businesscard /opt/businesscard

EXPOSE 8080

ENV NAME "Sergio"
ENTRYPOINT /opt/businesscard -name ${NAME}