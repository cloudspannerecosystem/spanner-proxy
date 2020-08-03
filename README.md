# spanner-proxy

An easy way to create Google Cloud Spanner proxies...

See examples/proxy for an example proxy server.

```go
proxy := proxy.New()
proxy.CreateSession = func(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
    // Your own session creation...
    return &pb.Session{
        Name:       "my-first-session",
        CreateTime: ptypes.TimestampNow(),
    }, nil
}
log.Fatal(proxy.Serve(lis))
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
