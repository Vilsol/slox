package slox

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestAll(t *testing.T) {
	var buf bytes.Buffer

	l := slog.New(slog.NewTextHandler(&buf, nil))

	ctx := context.Background()
	ctx = Into(ctx, l)
	ctx = With(ctx, "i-am", "outside", slog.Int("one", 1))
	ctx = WithGroup(ctx, "my-group")
	ctx = With(ctx, "hello", "world", slog.Int("two", 2))
	ctx = WithGroup(ctx, "another")

	testza.AssertFalse(t, Enabled(ctx, slog.LevelDebug))

	Debug(ctx, "message", "foo", "bar")
	testza.AssertEqual(t, buf.String(), "")
	buf.Reset()

	Info(ctx, "message", "foo", "bar")
	testza.AssertContains(t, buf.String(), "level=INFO msg=message i-am=outside one=1 my-group.hello=world my-group.two=2 my-group.another.foo=bar")
	buf.Reset()

	Warn(ctx, "message", "foo", "bar")
	testza.AssertContains(t, buf.String(), "level=WARN msg=message i-am=outside one=1 my-group.hello=world my-group.two=2 my-group.another.foo=bar")
	buf.Reset()

	Error(ctx, "message", "foo", "bar")
	testza.AssertContains(t, buf.String(), "level=ERROR msg=message i-am=outside one=1 my-group.hello=world my-group.two=2 my-group.another.foo=bar")
	buf.Reset()

	Log(ctx, slog.LevelInfo, "message")
	testza.AssertContains(t, buf.String(), "level=INFO msg=message i-am=outside one=1 my-group.hello=world my-group.two=2")
	buf.Reset()

	LogAttrs(ctx, slog.LevelInfo, "message", slog.Int("three", 3))
	testza.AssertContains(t, buf.String(), "level=INFO msg=message i-am=outside one=1 my-group.hello=world my-group.two=2 my-group.another.three=3")
	buf.Reset()
}

func TestNoContext(t *testing.T) {
	logger := From(nil) //nolint:staticcheck
	testza.AssertNotNil(t, logger)

	ctx := Into(nil, logger) //nolint:staticcheck
	testza.AssertNotNil(t, ctx)
}

func TestNoLogger(t *testing.T) {
	logger := From(context.Background())
	testza.AssertNotNil(t, logger)
}
