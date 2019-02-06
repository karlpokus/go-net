# go-net
Extending the [net.Conn interface](https://www.youtube.com/watch?v=afSiVelXDTQ&t=1226s)

# usage
server-client
```bash
# server
$ go run *.go
# client
$ nc localhost 37042
```

proxy
```bash
# server 1
$ go run *.go
# server 2
$ nc -l 37043
# client
$ nc localhost 37042
# data will be proxied to server 2
```

notes
- `l.Accept` and `io.Copy` are blocking
- `io.Copy` copies to dest.Write from src.Read until EOF
- Never read anything from the conn in the accept-loop. Pass it on to prevent total outage
- Always use timeouts to prevent hanging clients and resource waste

# todos
- [x] client-server
- [x] multiplexing
- [x] timeout
