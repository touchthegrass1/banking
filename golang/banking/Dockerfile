FROM golang:1.19 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 go build -o /go/bin/app
CMD ["/go/bin/app"]

# # Now copy it into our base image.
# FROM gcr.io/distroless/static-debian11
# COPY --from=build /go/bin/app /
# EXPOSE 8080/tcp
# CMD ["/app"]