# Airplane Go SDK ![CI status](https://img.shields.io/github/workflow/status/airplanedev/go-sdk/test/main) ![License](https://img.shields.io/github/license/airplanedev/go-sdk) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/airplanedev/go-sdk)

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
      Name   string `json:"name"`
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
