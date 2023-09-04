# slox

[![push](https://github.com/Vilsol/slox/actions/workflows/push.yml/badge.svg)](https://github.com/Vilsol/slox/actions/workflows/push.yml)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/vilsol/slox)
[![codecov](https://codecov.io/gh/Vilsol/slox/branch/master/graph/badge.svg?token=LFNKYWS0N2)](https://codecov.io/gh/Vilsol/slox)
[![CodeFactor](https://www.codefactor.io/repository/github/vilsol/slox/badge)](https://www.codefactor.io/repository/github/vilsol/slox)
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
