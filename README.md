# slox

[![push](https://github.com/Vilsol/slox/actions/workflows/push.yml/badge.svg)](https://github.com/Vilsol/slox/actions/workflows/push.yml)
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/vilsol/slox)
[![codecov](https://codecov.io/gh/Vilsol/slox/graph/badge.svg?token=633UR1LGRC)](https://codecov.io/gh/Vilsol/slox)
[![Go Report Card](https://goreportcard.com/badge/github.com/Vilsol/slox)](https://goreportcard.com/report/github.com/Vilsol/slox)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vilsol/slox)
[![Go Reference](https://pkg.go.dev/badge/github.com/Vilsol/slox.svg)](https://pkg.go.dev/github.com/Vilsol/slox)

A simple context wrapper around builtin Go [slog](https://pkg.go.dev/log/slog) library.

## Usage

```go
package main

import (
	"context"
	"log/slog"

	"github.com/Vilsol/slox"
)

func main() {
	ctx := context.Background()
	ctx = slox.With(ctx, "foo", "bar")
	ctx = slox.WithGroup(ctx, "my-group")
	ctx = slox.With(ctx, slog.Int("one", 1), "hello", "world")
	slox.Info(ctx, "message")
}
```
