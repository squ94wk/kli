# kli - the key cli
`kli` is a simple CLI to inspect and manage cryptographic keys, certificates and PKI for local development.

```shell script
# create new private key
kli key create ca.key

# create pki (root CA, ...) with new profile
kli pki bootstrap --interactive --new-profile

# show info about certificate in ca.crt
kli inspect ca.crt

# verify cert against PKI of profile
kli cert verify tls.crt

# print named cert "ca" to stdout in pem format
kli cert get -f pem ca

# create PKI according to yml file description, deploy as secrets
kli apply -f ./kli.yml
```