#this is a dockerfile for httpserver
#author: powerfj
FROM golang AS builder
ENV GOOS=linux
ENV GO111MODULE=off
ENV CGO_ENABLED=0
WORKDIR /build
COPY . ./
RUN go build -o httpserver httpserver.go

FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"]

