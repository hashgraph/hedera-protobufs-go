# Hedera™ Hashgraph Go Protubufs

> Generated protobufs in Go for interacting with Hedera Hashgraph: the official distributed
> consensus platform built using the hashgraph consensus algorithm for fast,
> fair and secure transactions. Hedera enables and empowers developers to
> build an entirely new class of decentralized applications.

## Install

```sh
$ go get github.com/hashgraph/hedera-protobufs-go
```

## Usage

```go
import "github.com/hashgraph/hedera-protobufs-go/services"
```

## Example

_to be written_

## Development

When updating the protobufs submodule, the generated code should be updated.

### Prerequisites

-   [Go](https://golang.org/doc/install) v15+

-   [Protobuf Compiler](https://developers.google.com/protocol-buffers)

-   Go plugins for the protobuf compiler

    ```sh
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

### Build

```sh
# update the proto/ directory from hedera-protobufs
$ git submodule update --init

# generate go code from protobuf definitions
$ go generate ./...

# ensure the projects build
$ go vet ./...
```

If you are experiencing generation difficulties:

-   Make sure all protobuf files in the proto subproject
    are at one folder level(no subdirs)
-   Remove protobuf import directory `state/` in the
    generated `.go` code.

## License

Licensed under Apache License,
Version 2.0 – see [LICENSE](LICENSE) in this repo
