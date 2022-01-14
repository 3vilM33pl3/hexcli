# Overview
Command line interface to manipulate game content

## Commands
### Hexagon manipulation
    nb hex add [0,0,0] [ref]
    nb hex del [0,0,0]
    nb hex update [0,0,0] [ref]
    nb hex info[0,0,0]
    nb hex get [0,0,0]
### Status hexagon network (storage server, meta data server, connected clients)
    nb status server
    nb status storage
    nb status clients
### Package content
    nb pack [dir]

## Run locally
Run [`hexcloud`](https://github.com/3vilM33pl3/hexcloud) server. The server listens to 0.0.0.0:8080 by default.

`go run .\cmd\hexcloud\hexcloud.go --local=true`

Test with status command:

`go run nb.go status server`
