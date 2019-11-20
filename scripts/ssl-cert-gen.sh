#!/bin/bash

SERVER_CN=localhost

# Generate Certificate Authority private key
openssl genrsa -passout pass:quality -des3 -out cert/ca.key 2048

# Generate Certificate Authority trust certificate
openssl req -passin pass:quality -new -x509 -days 365 -key cert/ca.key -out cert/ca.crt -subj "/CN=${SERVER_CN}"

# Generate Server private key
openssl genrsa -passout pass:quality -des3 -out cert/server.key 2048

# Generate a Certificate Signing Requests (CSR)
openssl req -passin pass:quality -new -key cert/server.key -out cert/server.csr -subj "/CN=${SERVER_CN}"

# Generate Server certificate signing the certificate with the CA
openssl x509 -req -passin pass:quality -days 365 -in cert/server.csr -CA cert/ca.crt -CAkey cert/ca.key -set_serial 01 -out cert/server.crt

# Convert Server private key to .pem format
openssl pkcs8 -topk8 -nocrypt -passin pass:quality -in cert/server.key -out cert/server.pem