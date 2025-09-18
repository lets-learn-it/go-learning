# SSH Tunnel

```sh
# start ssh servers.
# use already available public and private keys.
podman compose up

# Compile
go build ./cmd/main.go

# run
./main -port=2222 -username=pskp -command="ls -la /"
```
