package main

import (
	"fmt"
	"log"
	"net/http"
	"project/configs"
	"project/internal"
	"project/internal/logger"
)

func webSiteRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// リクエストログを追加
		logger.Info("Page accessed",
			"path", r.URL.Path,
			"method", r.Method,
			"user_agent", r.UserAgent(),
			"remote_addr", r.RemoteAddr,
		)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
}

func main() {
	fmt.Println("Web Site Main")

	// 1. 設定読み込み
	config, err := configs.LoadConfig("configs/config-site.toml")
	if err != nil {
		// loggerが初期化されていないので、標準のlogを使用
		log.Fatal("設定ファイルの読み込みに失敗しました:", err)
	}

	// 2. logger初期化（設定読み込み後）
	logger.InitLogger(logger.LoggerLevel(config.Log.Level), config.App.Environment, config.Log.File)

	// 3. WEBサイト開始ログ
	logger.Info("Application starting",
		"name", config.App.Name,
		"version", config.App.Version,
		"environment", config.App.Environment,
		"port", config.Server.Port,
	)

	// 4. サーバー起動
	serverConfig := internal.ServerConfig{
		Port:   config.Server.Port,
		Routes: webSiteRoutes,
	}

	server := serverConfig.SetupServer()
	if err := internal.StartServerWithGracefulShutdown(server); err != nil {
		logger.Error("Server failed", "error", err)
	}
}
