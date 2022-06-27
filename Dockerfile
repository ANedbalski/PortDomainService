FROM golang:1.17 AS build

COPY . /go/src/git@github.com/ANedbalski/PortDomainService.git

RUN cd /go/src/git@github.com/ANedbalski/PortDomainService.git \
    && GOOS=linux GOARCH=amd64 go install \
        -mod vendor \
        gitlab.com/bottle/core/customer-statements/cmd/...

FROM debian:latest

# Update CA certs
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy over app binary
COPY --from=build /go/bin/* /usr/sbin/


# Add a user
RUN adduser --disabled-login app
USER app

WORKDIR /app/

CMD [ "/usr/sbin/ports", "run" ]
