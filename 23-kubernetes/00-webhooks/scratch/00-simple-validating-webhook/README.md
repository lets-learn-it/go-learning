# Simple Validating Webhook

## Create Certs

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -addext "subjectAltName = DNS:my-service.default.svc" -keyout ca.key -out ca.crt

# caBundle in webhook.yaml
cat ca.crt | base64 | fold
```

## References

[[00] A Simple Kubernetes Admission Webhook - slack.engineering](https://slack.engineering/simple-kubernetes-webhook/)

[[01] Writing Validating Webhook Controller From Scratch - 1 | Validating Webhook Kubernetes - youtube.com](https://www.youtube.com/watch?v=NRmNE-6Wb7g)