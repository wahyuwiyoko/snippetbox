# Snippetbox

Simple web application that used to paste and share snippets of text.

## Installation

For production server, it is recommend using [Let's Encrypt](https://letsencrypt.org/)
to create TLS certificate. But for development purpose, it can generate
self-signed TLS certificate with Go's standard library:

```sh
mkdir tls
cd tls
go run {your-go-path}/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Run the web server:

```sh
go run cmd/web/!(*_test).go
```
