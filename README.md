# spanner-proxy

An easy way to write Google Cloud Spanner proxies...



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