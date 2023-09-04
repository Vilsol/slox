package slox

import (
	"context"
	"log/slog"
)

type slogKey struct{}

// Into wraps the provided logger in a context
func Into(ctx context.Context, logger *slog.Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, slogKey{}, logger)
}

// From returns the logger from the context
//
// If the context does not contain a logger, a default logger is returned
func From(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return slog.Default()
	}

	if logger := ctx.Value(slogKey{}); logger != nil {
		return logger.(*slog.Logger)
	}

	return slog.Default()
}

// With returns a context.Context that includes the given attributes
// in each output operation. Arguments are converted to
// attributes as if by [slog.Logger.Log].
//
// Internally calls [slog.Logger.With]
func With(ctx context.Context, args ...any) context.Context {
	return Into(ctx, From(ctx).With(args...))
}

// WithGroup returns a context.Context that starts a group, if name is non-empty.
// The keys of all attributes added to the Logger will be qualified by the given
// name. (How that qualification happens depends on the [slog.Handler.WithGroup]
// method of the Logger's Handler.)
//
// If name is empty, WithGroup returns the receiver.
//
// Internally calls [slog.Logger.WithGroup]
func WithGroup(ctx context.Context, name string) context.Context {
	return Into(ctx, From(ctx).WithGroup(name))
}

// Enabled reports whether context.Context emits log records at the given level.
//
// Internally calls [slog.Logger.Enabled]
func Enabled(ctx context.Context, level slog.Level) bool {
	return From(ctx).Enabled(ctx, level)
}

// Log emits a log record with the current time and the given level and message.
// The Record's Attrs consist of the Logger's attributes followed by
// the Attrs specified by args.
//
// The attribute arguments are processed as follows:
//   - If an argument is an Attr, it is used as is.
//   - If an argument is a string and this is not the last argument,
//     the following argument is treated as the value and the two are combined
//     into an Attr.
//   - Otherwise, the argument is treated as a value with key "!BADKEY".
//
// Internally calls [slog.Logger.Log]
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	From(ctx).Log(ctx, level, msg, args...)
}

// LogAttrs is a more efficient version of [slog.Logger.Log] that accepts only Attrs.
//
// Internally calls [slog.Logger.LogAttrs]
func LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	From(ctx).LogAttrs(ctx, level, msg, attrs...)
}

// Debug logs at slog.LevelDebug.
//
// Internally calls [slog.Logger.DebugContext]
func Debug(ctx context.Context, msg string, args ...any) {
	From(ctx).DebugContext(ctx, msg, args...)
}

// Info logs at slog.LevelInfo.
//
// Internally calls [slog.Logger.InfoContext]
func Info(ctx context.Context, msg string, args ...any) {
	From(ctx).InfoContext(ctx, msg, args...)
}

// Warn logs at slog.LevelWarn.
//
// Internally calls [slog.Logger.WarnContext]
func Warn(ctx context.Context, msg string, args ...any) {
	From(ctx).WarnContext(ctx, msg, args...)
}

// Error logs at slog.LevelError.
//
// Internally calls [slog.Logger.ErrorContext]
func Error(ctx context.Context, msg string, args ...any) {
	From(ctx).ErrorContext(ctx, msg, args...)
}
