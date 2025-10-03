package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"project/configs"
	"project/internal/logger"
	"syscall"
	"time"
)

type ServerConfig struct {
	Port    int
	Routes  func(*http.ServeMux)
	DevMode bool
}

func disableCacheInDevMode(devMode bool, next http.Handler) http.Handler {
	if !devMode {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (config ServerConfig) SetupServer() (*http.Server, error) {
	err := configs.GetConfig().Validate()
	if err != nil {
		logger.Error("Invalid server config: %v", err)
		return nil, fmt.Errorf("invalid server config: %w", err)
	}

	mux := http.NewServeMux()

	// 静的ファイルを提供
	mux.Handle("/assets/", disableCacheInDevMode(config.DevMode, http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))))


	if config.Routes != nil {
		config.Routes(mux)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: mux,
	}

	return server, nil
}

// サーバーを起動し、シグナル待機してグレースフルシャットダウンを行う
// シグナル待機は、Ctrl+Cなどのシグナルを待ち、グレースフルシャットダウンを行う
// グレースフルシャットダウンは、サーバーを停止する前に、現在のリクエストが完了するのを待つ
func StartServerWithGracefulShutdown(server *http.Server) error {
	// サーバーを別のgoroutineで起動
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// シグナル待機（Ctrl+Cなど）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// 30秒以内にグレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown:", err)
	}
	logger.Info("Server exited")
	return nil
}
