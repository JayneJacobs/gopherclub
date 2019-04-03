FROM golang:1.11.5-alpine3.9
MAINTAINER Jayne Jacobs jaynejacobs@jaynejacobs.com

# Declare required environment variables
ENV gopherclub_APP_ROOT=/go/src/github.com/JayneJacobs/gopherclub
ENV gopherclub_HASH_KEY="CRKVBJs0kfyeQ9Y1"
ENV gopherclub_BLOCK_KEY="9LtmRLzVH27CwxrO"
ENV GOPATH=/go

# Get the required Go packages
RUN go get -u github.com/gopherjs/gopherjs
RUN go get -u honnef.co/go/js/dom
RUN go get -u -d -tags=js github.com/gopherjs/jsbuiltin
RUN go get -u honnef.co/go/js/xhr
RUN go get -u github.com/gopherjs/websocket
RUN go get -u go.isomorphicgo.org/go/isokit 
RUN go get -u github.com/JayneJacobs/gopherclub/tdewolff/minify@v2.3.4
RUN go get -u github.com/JayneJacobs/gopherclub
RUN go get -u go.go.org/uxtoolkit/cog
RUN go get  github.com/JayneJacobs/gopherclub/vendor/tdewolff/parse@v2.3.6


# Transpile and install the client-side application code
RUN cd $gopherclub_APP_ROOT/client; go get ./..; /go/bin/gopherjs build -m --verbose -o $gopherclub_APP_ROOT/static/js/client.min.js

# Build and install the server-side application code
RUN go install github.com/JayneJacobs/gopherclub

# Specify the entrypoint
ENTRYPOINT /go/bin/gopherclub

# Expose port 8080 of the container
EXPOSE 8081
