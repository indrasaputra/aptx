# Frequently Asked Questions

1. `make mockgen` returns `imported package collision: "backoff" imported twice`

    `mockgen` must use [reflect mode](https://github.com/golang/mock#reflect-mode) to understand the interface.
    This error will likely to happen if we want to generate mock for interface in generated `*.pb.go` files.
    For example, we want to generate mock for `AptxService_StreamAllURLServer` in `proto/indrasaputra/aptx/v1/aptx_grpc.pb.go`.
    We will get that reported error. To resolve this issue, run mockgen in [reflect mode](https://github.com/golang/mock#reflect-mode).
    ```
    $ cd proto/indrasaputra/aptx/v1
    $ mockgen -destination=<your-destination.go> . AptxService_StreamAllURLServer
    ```
    