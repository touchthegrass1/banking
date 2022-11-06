FROM golang:1.19 AS build
WORKDIR /go/src
COPY . ./banking

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o bankingapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/bankingapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./bankingapi"]
