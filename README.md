# Self Signed Certs

A bit of research into self signed certificates, and using them with Golang and
Electron.

**Editor's Notes**:

- [Christine](github.com/Xe) recommends
  [minica](https://github.com/jsha/minica) for managing certs which probably
  isn't a bad idea. This example uses the OpenSSL CLI though for learning.
- Go examples were borrowed from:
  - [Server](https://gist.github.com/denji/12b3a568f092ab951456)
  - [Client](https://stackoverflow.com/questions/38822764/how-to-send-a-https-request-with-a-certificate-golang)

## Generating a RootCA

To generate a root key

```
openssl genrsa     \
  -out rootCA.key  \
  4096
```

To then generate a CRT to go along with it  
(This creates a public cert that last 10 years)

```
openssl req        \
  -x509            \
  -new             \
  -key rootCA.key  \
  -days 3650       \
  -out rootCA.crt  \
```

The root CA shouldn't be used when serving a resource. Instead issue a cert and
key for the server to use.

## Issuing a Cert for the Server

Generate a new key

```
openssl genrsa      \
  -out $DOMAIN.key  \
  2048
```

Generate a CSR to help generate our Cert file

```
openssl req         \
  -new              \
  -key $DOMAIN.key  \
  -out $DOMAIN.csr
```

Issue our certificate

```
openssl x509         \
  -req               \
  -in $DOMAIN.csr    \
  -CA rootCA.crt     \
  -CAkey rootCA.key  \
  -CAcreateserial    \
  -days 365          \
  -out $DOMAIN.crt
```

When all is said and done, we should now have:

```
.
├── $DOMAIN.crt
├── $DOMAIN.csr
├── $DOMAIN.key
├── rootCA.crt
├── rootCA.key
└── rootCA.srl
```

## Clients and Servers

Give the client the `$DOMAIN.crt` and throw the server the `$DOMAIN.key` and
`$DOMAIN.crt`, and protect the `rootCA.key`.
