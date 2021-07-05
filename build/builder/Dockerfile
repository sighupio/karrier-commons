# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# hadolint ignore=DL3007
FROM quay.io/sighup/policeman:latest as linter

ENV VALIDATE_KUBERNETES_KUBEVAL="false"

RUN mkdir /app
WORKDIR /app

COPY . .

RUN /entrypoint.sh

FROM golang:1.16 as license

RUN go get -u github.com/google/addlicense

RUN mkdir /app
WORKDIR /app

COPY . .

RUN addlicense -c "SIGHUP s.r.l" -v -l bsd --check .

FROM golang:1.16 as tester

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY pkg pkg
COPY internal internal

RUN go test -v ./... -cover

FROM golang:1.16 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY pkg pkg
COPY internal internal

WORKDIR /app/cmd/fip-commons
RUN go build

FROM debian:buster as release

COPY --from=builder /app/cmd/fip-commons/fip-commons /fip-commons

RUN groupadd -r SIGHUP && \
    useradd -u 1001 -r -s /bin/false -g SIGHUP cloud && \
    chown cloud:SIGHUP /fip-commons

USER 1001

ENTRYPOINT ["/fip-commons"]
