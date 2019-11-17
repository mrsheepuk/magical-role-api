FROM golang:latest

WORKDIR /go/src/github.com/mrsheepuk/magicalroleapi

# Add just the go.mod and go.sum before go mod download so we don't have to
# redownload if only source code has changed.
ADD ./go.mod .
ADD ./go.sum .
RUN go mod download

# Add the rest of the code and build
ADD . .
RUN go install github.com/mrsheepuk/magicalroleapi/cmd/magicalroleapi

ENTRYPOINT /go/bin/magicalroleapi
EXPOSE 8080
