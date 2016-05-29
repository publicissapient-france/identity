FROM alpine

RUN apk add --update ca-certificates # Certificates for SSL

COPY static /identity/static
COPY identity.html /identity
COPY identity /identity

WORKDIR /identity

EXPOSE 8080

ENTRYPOINT ["/identity/identity"]