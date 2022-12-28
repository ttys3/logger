package tests

import (
	"github.com/ttys3/slogsimple"
	"net"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestSlogLogging(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestSlogWith(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr)))

	l := slog.With("name", "Al")
	l.Info("hello", "age", 18)
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestSlogCustomOptions(t *testing.T) {
	th := slogsimple.NewHandlerOptions(slog.LevelInfo).NewJSONHandler(os.Stderr)
	slog.SetDefault(slog.New(th))

	l := slog.With("name", "Al")
	l.Info("hello", "age", 18)
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.Warn("this is warning")
	slog.Debug("this debug message should not shown up")
}

func TestSlogWithAtomicLevelVar(t *testing.T) {
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelInfo)
	th := slogsimple.NewHandlerOptions(lvl).NewJSONHandler(os.Stderr)
	slog.SetDefault(slog.New(th))

	l := slog.With("name", "Al")
	l.Info("hello", "age", 18)
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.Warn("this is warning")
	slog.Debug("this debug message should NOT shown up")
	lvl.Set(slog.LevelDebug)
	slog.Debug("this debug message should shown up")
}

func TestNewLogHandler(t *testing.T) {
	slogsimple.InitDefault()

	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.Debug("this debug message should NOT shown up")
}
