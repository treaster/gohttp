#!/bin/bash

# Generate self-signed TLS/SSL certs suitable for use (only) with gohttp server
# defined in this repo.

certname=localhost
if [ $# == 0 ]; then
    echo 'Using default certname "$certname". Override with "$0 [certname]".'
elif [ $# == 1 ]; then
    certname="$1"; shift
else
    echo 'Usage: $0 [optional certname]'
    exit 1
fi

openssl req \
    -new \
    -newkey rsa:2048 \
    -nodes \
    -keyout "${certname}.key" \
    -subj "/CN=${certname}" \
    -out "${certname}.csr"

openssl x509 -req \
    -days 365 \
    -in "${certname}.csr" \
    -signkey "${certname}.key" \
    -out "${certname}.crt"
