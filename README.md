# SanPush CLI

A Simple CLI for Deno Webhook Server to Reload and Send Message to Clients from Terminal.  

## Usage

- Download or clone the repo

```sh
https://github.com/sanwebinfo/sanpush-cli.git
cd sanpush-cli
```

- Create `yaml` File for Store API URL and Auth KEY

```sh

mkdir -p $HOME/sanpush
nano $HOME/sanpush/config.yaml

```

```yaml

bearer_token: your_secret_api_key
api_url: http://localhost:8000

```

- Execute the Script

```sh
go run sanpush.og -h
```

## Build Package

- Run Make file to build a package for your Systems

```sh
make build
```

## Packges Build for  

Linux, Apple, Windows and Android - `/makefile`  

- Linux-386
- Linux-arm-7
- Linux-amd64
- Linux-arm64
- Andriod-arm64
- windows-386
- windows-amd64
- darwin-amd64
- darwin-arm64

```sh
chmod +x sanpush
./sanpush -h
```

## API

Deno Webhook Server is a real-time server application built using Deno. It supports WebSocket connections for real-time updates it includes endpoints for triggering reloads and retrieving messages - **<https://github.com/sanwebinfo/deno-webhook-server>**

## LICENSE

MIT
