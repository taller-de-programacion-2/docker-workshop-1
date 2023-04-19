FROM golang:1.18-alpine

# Declare basic env vars
ENV GOPATH /go/
ENV DUMMY_APP_PATH /go/src/dummy-backend-service/

# Set up project directory on container
WORKDIR ${DUMMY_APP_PATH}
COPY . ${DUMMY_APP_PATH}

# Build project
RUN apk update && apk add bash make
RUN make build

ENTRYPOINT ["./build/dummy-backend"]
