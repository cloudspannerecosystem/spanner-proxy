# spanner-proxy

An easy way to create Google Cloud Spanner proxies...

See examples/proxy for an example proxy server.

```go
p := proxy.New()
p.CreateSession = func(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
    // Your own session creation...
    return &pb.Session{
        Name:       "my-first-session",
        CreateTime: ptypes.TimestampNow(),
    }, nil
}
log.Fatal(p.Serve(lis))
```

Proxy is honored by the official Cloud Spanner libraries automatically
if you set SPANNER_EMULATOR_HOST environment variable to the proxy endpoint.

```bash
export SPANNER_EMULATOR_HOST=localhost:9777
go run examples/client/main.go
```

Alternatively, you can use the Cloud Spanner gRPC clients
with the proxy endpoint and insecure connection:

```go
conn, err := grpc.Dial("localhost:9777", grpc.WithInsecure())
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

## When to write a proxy server?

There are various reasons you might want to write a proxy server.
Some examples:

* To override the behavior of the Cloud Spanner API.
* To manage connections and sessions outside of application server. 
* To instrument the Cloud Spanner calls in a custom way for monitoring
  and debugging purposes.
* To handle the authentication to Cloud Spanner outside of the application servers.