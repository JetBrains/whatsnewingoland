# syntax = docker/dockerfile:1.2
FROM golang:1.16.2-alpine3.13 as build

# Step 1. Observe that you can build a stage independently now with a click on the 🔨 icon next to each stage.

RUN apk --no-cache add git

WORKDIR /

RUN git clone https://github.com/go-delve/delve.git

WORKDIR /delve

RUN --mount=type=cache,target=/root/.cache/go-build go build -mod=vendor -o /dlv ./cmd/dlv

FROM alpine:3.13.2

# Step 2. Observe that you get code completion for stage name in the copy command below.
COPY --from=build /dlv /dlv

CMD ["/dlv", "version"]
