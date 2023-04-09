#!/bin/bash

# creating public/private key pairs
openssl genpkey -algorithm RSA -out ./keys/jwt.key.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -in ./keys/jwt.key.pem -pubout -out ./keys/jwt.pub.pem

# certificate, optional
openssl req -x509 -new -nodes -key ./keys/jwt.key.pem -sha256 -out ./keys/jwt.cert.pem -subj "/CN=unused" -days 36500