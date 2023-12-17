package provider

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, nil)
	l := slog.New(h)
	return l
}
