FROM golang:1.18-alpine AS build
WORKDIR /go/src/github.com/milQA/tzstat
RUN apk add make

COPY . .
RUN make build

FROM alpine:3.10 AS release
COPY --from=build /go/src/github.com/milQA/tzstat/bin/tzstat /usr/local/bin/tzstat

RUN apk add --no-cache make bash curl tzdata
ENTRYPOINT ["/usr/local/bin/tzstat"]
