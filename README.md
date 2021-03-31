# Airplane Go SDK [![Docs](https://img.shields.io/github/v/tag/airplanedev/go-sdk?label=docs)](https://pkg.go.dev/github.com/airplanedev/go-sdk) [![License](https://img.shields.io/github/license/airplanedev/go-sdk)](https://github.com/airplanedev/go-sdk/blob/main/LICENSE) [![CI status](https://img.shields.io/github/workflow/status/airplanedev/go-sdk/test/main)](https://github.com/airplanedev/go-sdk/actions?query=branch%3Amain)

Go SDK for writing Airplane.dev tasks.

## Getting started

```sh
go get github.com/airplanedev/go-sdk
```

## Usage

```go
package main

import (
  "context"
  "fmt"

  airplane "github.com/airplanedev/go-sdk"
)

func main() {
  airplane.Run(func(ctx context.Context) error {
    var parameters struct {
      Name string `json:"name"`
    }
    if err := airplane.Parameters(&parameters); err != nil {
      return err
    }

    msg := fmt.Sprintf("Hello, %s!\n", parameters.Name)
    airplane.MustOutput(msg)

    return nil
  })
}
```
