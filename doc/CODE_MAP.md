# Code Map

Code Map explains how the codes are structured in this project. This document will list down all folders or packages and their purpose.

## `.github`

This folder is used to define the [Github Actions](https://docs.github.com/en/actions).

## `bin`

This folder contains any executable binary to support the project.
For example, there is `generate-grpc.sh`. It is a shell script file used to generate gRPC related code from Protocol Buffer definition.

## `cmd`

This folder contains the `main.go`.
The use case may be run and served in multi forms, such as API, cron, or fullstack web.
To cater that case, `cmd` folder can contains subfolders with each folder named accordingly to the form and contain only main package.
e.g: `cmd/api/main.go`, `cmd/cron/main.go`, and `cmd/web/main.go`

For this project, we prefer to use `cmd/server/main.go` and `cmd/client/main.go` as our use cases are only in the form of gRPC server and client.

## `doc`

This folder contains all documentations related to the project.

## `entity`

This folder contains the domain of the application.
Mostly, this folder contains only structs, constants, global variables, enumerations, or functions with simple logic related to the core domain of the application (not a business logic).
Since we use Protocol Buffer, entity has a close (tightly coupled) relationship with any struct generated from `.proto` files.

## `infrastructure`

This folder contains any configuration for deployment infrastructure, such as monitoring, logging, kubernetes, etc.

## `proto`

This folder contains `.proto` files and all files generated from or based on `.proto`.

## `openapiv2`

This folder contains API definition for HTTP/1.1 REST.
The contents of this folder are generated from `.proto` as well.

## `usecase`

This folder contains the main business logic of the project.
All interfaces and the business logic flows are defined here.
If someone wants to know the flow of the application, they better start to open this folder.

## `test`

This folder contains test related stuffs.
For the case of unit test, the unit test files are put in the same directory as the files that they test.
It is one of the Go best practice, so we follow.

### `test/fixture`

This folder contains a well defined support for test.

### `test/mock`

This folder contains mock for testing.

## `internal`

All APIs/codes in the internal folder (and all if its subfolders) are designed to [not be able to be imported](https://golang.org/doc/go1.4#internalpackages).
This folder contains all detail implementation specified in the `usecase` folder.

### `internal/builder`

This folder contains the [builder design pattern](https://sourcemaking.com/design_patterns/builder).
It composes all codes needed to build a full usecase.

### `internal/config`

This folder contains configuration for the application.

### `internal/http2/grpc/handler`

This folder contains the HTTP/2 gRPC handlers.
Codes in this folder implement gRPC server interface located in `proto` directory.

### `internal/repository`

This folder contains codes that connect to the repository, such as database.
Repository is not limited to databases. Anything that can be a repository can be put here.

### `internal/tool`

This folder contains all codes that can support code the project.