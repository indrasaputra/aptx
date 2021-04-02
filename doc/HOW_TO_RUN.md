# How to Run

- Create `.env` file

    You can copy the `env.sample` and change its values.

    ```
    $ cp env.sample .env
    ```

- Fill `PORT_GRPC` and `PORT_HTTP` value as you wish.
    `PORT_GRPC` is a port for HTTP/2 gRPC. `PORT_HTTP` is port for HTTP/1.1 REST.
    We encourage to let both values as default.

- Download the dependencies

    ```
    $ make tidy
    ```

- Run the application

    ```
    $ go run cmd/server/main.go
    ```