#!/bin/bash

cd keys
rm -rf *.pem *.key *.crt *.csr

# -------------------------------
# Generate CA key and certificate
# -------------------------------

openssl genrsa -out ca-key.pem
openssl req -new -x509 -days 365 -key ca-key.pem -out ca-cert.pem -subj "/C=PT/CN=CA"


# --------------------------------------------------------------------------
# Generate server key and certificate and let CA sign the server certificate
# --------------------------------------------------------------------------

openssl genrsa -out server-key.pem
openssl req -new -key server-key.pem -out server-csr.pem -subj "/C=PT/ST=Lisbon/L=Lisbon/O=Ride/CN=localhost"
openssl x509 -req -days 120 -in server-csr.pem -CA ca-cert.pem -CAkey ca-key.pem -extfile server-ext.cnf -set_serial 01 -out server-cert.pem

rm server-csr.pem
rm ca-key.pem

echo ""
echo "done!"