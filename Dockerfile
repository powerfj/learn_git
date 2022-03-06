#this is a dockerfile for httpserver
#author: powerfj
#第一行指定基础镜像
FROM golang
#镜像的操作指令
WORKDIR $GOPATH/src/httpserver
COPY . ./
RUN go build httpserver.go
EXPOSE 7070
ENTRYPOINT ["./httpserver"]

