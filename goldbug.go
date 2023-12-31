package goldbug

import (
	"log/slog"
	"os"
	"path/filepath"
)

func setPartialPath(source *slog.Source) {
	fileName := filepath.Base(source.File)
	parentDir := filepath.Base(filepath.Dir(source.File))

	source.File = filepath.Join(parentDir, fileName)
}

func SetDefaultLoggerText(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					setPartialPath(source)
				}
			}
			return a
		},
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)
	
	slog.SetDefault(logger)
}

func SetDefaultLoggerJson(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					setPartialPath(source)
				}
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stderr, &opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}
