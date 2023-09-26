# Let's Go 
This is a follow along example of the book `Let's Go`. 

### Generate TLS Self-signed Certificate
```
mkdir tls
cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
