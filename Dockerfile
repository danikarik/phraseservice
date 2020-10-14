# Simple usage with a mounted data directory:
# > docker build -t phraseservice .
#
# Server:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.phraseservice:/root/.phraseservice phraseservice phraseserviced init test-chain
# TODO: need to set validator in genesis so start runs
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.phraseservice:/root/.phraseservice phraseservice phraseserviced start
#
# Client: (Note the phraseservice binary always looks at ~/.phraseservice we can bind to different local storage)
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.phraseservicecli:/root/.phraseservice phraseservice phraseserviced keys add foo
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.phraseservicecli:/root/.phraseservice phraseservice phraseserviced keys list

FROM golang:alpine AS build-env

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

WORKDIR /go/src/app

COPY . .

RUN make build-linux

FROM alpine:edge

RUN apk add --update ca-certificates
WORKDIR /root

COPY --from=build-env /go/src/app/build/phraseserviced /usr/bin/phraseserviced

EXPOSE 26656 26657 1317 9090

CMD ["phraseserviced"]
