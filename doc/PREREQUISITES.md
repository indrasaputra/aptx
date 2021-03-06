# Prerequisites

- Install Go

    We use version 1.16. Follow [Golang installation guideline](https://golang.org/doc/install).

- Install golangci-lint

    Follow [golangci-lint installation](https://golangci-lint.run/usage/install/).

- Install gomock

    Follow [gomock installation](https://github.com/golang/mock).

- Install Buf

    Follow [Buf installation](https://docs.buf.build/installation).

- Install clang-format

    We use [clang-format](https://clang.llvm.org/docs/ClangFormat.html) to format our protocol buffer files.
    We use version 11.1.0 when we develop this project.

- Install `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`, `protoc-gen-openapiv2`, `protoc-gen-go` 

    ```
    $ go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

    That will place four binaries in $GOBIN;

    - `protoc-gen-go-grpc`
    - `protoc-gen-grpc-gateway`
    - `protoc-gen-openapiv2`
    - `protoc-gen-go`

    Make sure that $GOBIN is in $PATH.

    For more this section installation guide, please refer to [grpc-gateway installation](https://github.com/grpc-ecosystem/grpc-gateway#installation).