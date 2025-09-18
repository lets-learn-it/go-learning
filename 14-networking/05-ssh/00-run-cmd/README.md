# Run Command

```sh
# start ssh server.
# use your own public key
podman run -e PUBLIC_KEY="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIERqgXxcSQR3BVH3E/xiz6hMbCPF+VYhXZ+B52U8QDLz pskp@mint" -p 2222:2222 linuxserver/openssh-server

# Compile
go build ./cmd/main.go

# run
./main -port=2222 -username=linuxserver.io -command="ls -la /"
```
