package main

import (
	"fmt"
	"log"
	"net/http"
	"project/configs"
	"project/internal"
	"project/internal/db"
	"project/internal/logger"
	"project/internal/web/app/company"
	"project/internal/web/app/user"
)

// グローバル変数でデータベース接続を管理
var database *db.DB

func webAppRoutes(mux *http.ServeMux) {
	user.Routes(mux, database)
	company.Routes(mux, database)

	mux.HandleFunc("/app/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	// データベースヘルスチェックエンドポイント
	mux.HandleFunc("/app/health", func(w http.ResponseWriter, r *http.Request) {
		if database == nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database not connected"))
			return
		}

		ctx := r.Context()
		if err := database.HealthCheck(ctx); err != nil {
			logger.Error("Database health check failed", "error", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database health check failed"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Database is healthy"))
	})
}

func main() {
	fmt.Println("Web App Main")

	// 1. 設定読み込み
	config, err := configs.LoadConfig("configs/config-app.toml")
	if err != nil {
		// loggerが初期化されていないので、標準のlogを使用
		log.Fatal("設定ファイルの読み込みに失敗しました:", err)
	}

	// 2. logger初期化（設定読み込み後）
	logger.InitLogger(logger.LoggerLevel(config.Log.Level), config.App.Environment, config.Log.File)

	// 3. アプリケーション開始ログ
	// logger.Info("Application starting",
	// 	"name", config.App.Name,
	// 	"version", config.App.Version,
	// 	"environment", config.App.Environment,
	// 	"port", config.Server.Port,
	// )

	// 4. データベース接続
	logger.Info("データベースに接続中...")
	database, err = db.NewDB(config)
	if err != nil {
		logger.Error("データベース接続に失敗しました", "error", err)
		log.Fatal("データベース接続に失敗しました:", err)
	}
	defer database.Close()

	logger.Info("データベース接続が完了しました")

	// 5. サーバー起動
	serverConfig := internal.ServerConfig{
		Port:   config.Server.Port,
		Routes: webAppRoutes,
	}

	server := serverConfig.SetupServer()
	if err := internal.StartServerWithGracefulShutdown(server); err != nil {
		logger.Error("Server failed", "error", err)
	}
}
