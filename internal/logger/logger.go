package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
)

// Logger ログレベル
type LoggerLevel string

const (
	DebugLevel LoggerLevel = "debug"
	InfoLevel  LoggerLevel = "info"
	WarnLevel  LoggerLevel = "warn"
	ErrorLevel LoggerLevel = "error"
)

// Logger アプリケーション全体で使用するロガー
var Logger *slog.Logger

// multiHandler（ターミナルとファイルなど） 複数のハンドラーに同時にログを出力するハンドラー
type multiHandler struct {
	handlers []slog.Handler
}

// Enabled いずれかのハンドラーが有効な場合にtrueを返す
func (m *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle すべてのハンドラーにログを出力
func (m *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		if h.Enabled(ctx, r.Level) {
			if err := h.Handle(ctx, r); err != nil {
				return err
			}
		}
	}
	return nil
}

// WithAttrs 属性を追加した新しいハンドラーを作成
func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &multiHandler{handlers: newHandlers}
}

// WithGroup グループを追加した新しいハンドラーを作成
func (m *multiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: newHandlers}
}

// InitLogger ロガーを初期化
func InitLogger(level LoggerLevel, environment string, logFile string) {
	var logLevel slog.Level

	switch level {
	case DebugLevel:
		logLevel = slog.LevelDebug
	case InfoLevel:
		logLevel = slog.LevelInfo
	case WarnLevel:
		logLevel = slog.LevelWarn
	case ErrorLevel:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	// ターミナル用のハンドラー（設定されたレベルで出力）
	terminalHandler := slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level:     logLevel, // 設定されたレベル（通常はInfo）
			AddSource: true,
		})

	// ファイル用のハンドラー（Errorレベル以上のみ）
	var fileHandler slog.Handler
	if logFile != "" {
		// ログディレクトリを作成
		logDir := filepath.Dir(logFile)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			// ディレクトリ作成に失敗した場合はファイルハンドラーをnilに
			slog.Error("ログディレクトリの作成に失敗しました", "error", err, "path", logDir)
			fileHandler = nil
		} else {
			// ログファイルを開く
			file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				// ファイルオープンに失敗した場合はファイルハンドラーをnilに
				slog.Error("ログファイルのオープンに失敗しました", "error", err, "file", logFile)
				fileHandler = nil
			} else {
				// ファイル用はErrorレベル以上のみ
				fileHandler = slog.NewTextHandler(file,
					&slog.HandlerOptions{
						Level:     slog.LevelError, // ファイルはErrorレベル以上のみ
						AddSource: true,
					})
			}
		}
	}

	// 複数のハンドラーを組み合わせる
	var handler slog.Handler
	if fileHandler != nil {
		// ファイルとターミナルの両方に出力
		handler = &multiHandler{
			handlers: []slog.Handler{terminalHandler, fileHandler},
		}
	} else {
		// ターミナルのみ
		handler = terminalHandler
	}

	// ロガーを作成(アプリケーション全体で使用するロガー)
	Logger = slog.New(handler).With(
		"service", "web-site",
		"environment", environment,
	)
}

// GetLogger ロガーを取得
func GetLogger() *slog.Logger {
	if Logger == nil {
		// デフォルト設定で初期化
		InitLogger(InfoLevel, "development", "logs/app.log")
	}
	return Logger
}

// 便利な関数を提供
func Debug(msg string, args ...any) {
	GetLogger().Debug(msg, args...)
}

func Info(msg string, args ...any) {
	GetLogger().Info(msg, args...)
}

func Warn(msg string, args ...any) {
	GetLogger().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	GetLogger().Error(msg, args...)
}

// コンテキスト付きのログ関数
func DebugContext(ctx context.Context, msg string, args ...any) {
	GetLogger().DebugContext(ctx, msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	GetLogger().InfoContext(ctx, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	GetLogger().WarnContext(ctx, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	GetLogger().ErrorContext(ctx, msg, args...)
}

// グループ化されたログ関数
func WithGroup(name string) *slog.Logger {
	return GetLogger().WithGroup(name)
}

// フィールド付きのログ関数
func With(args ...any) *slog.Logger {
	return GetLogger().With(args...)
}
