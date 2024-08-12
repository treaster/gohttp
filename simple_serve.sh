#!/bin/bash

if [ $# != 1 ]; then
    echo "usage: $(basename $0) [directory to serve]"
    exit 1
fi

relative_script_dir="$(dirname $0)"
cert_dir=/tmp/gohttp

relative_serve_dir=$1; shift
full_serve_dir="$(pwd)/${relative_serve_dir}"
full_script_dir="$(pwd)/${relative_script_dir}"

cd "${full_script_dir}"

if [ ! -f "${cert_dir}/localhost.crt" ]; then
    echo "generating new certs in ${cert_dir}..."
    mkdir -p "${cert_dir}"
    pushd "${cert_dir}"
    "${full_script_dir}/generate_tls_certs.sh"
    popd
fi

go run . \
    --port=4443 \
    --certname="${cert_dir}/localhost" \
    --dir="${full_serve_dir}"
