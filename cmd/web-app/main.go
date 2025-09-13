package main

import (
	"fmt"
	"log"
	"project/configs"
	"project/internal"
	"project/internal/db"
	"project/internal/logger"
	"project/internal/web/app"
)

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

	// 3. データベース接続
	database, err := db.NewDB(config)
	if err != nil {
		logger.Error("データベース接続に失敗しました", "error", err)
		log.Fatal("データベース接続に失敗しました:", err)
	}
	defer database.Close()

	// 4. アプリケーションルートを設定
	appRoutes := app.Routes(database)

	// 5. サーバー起動
	serverConfig := internal.ServerConfig{
		Port:    config.Server.Port,
		Routes:  appRoutes,
		DevMode: config.App.Environment == "development",
	}

	server, err := serverConfig.SetupServer()
	if err != nil {
		logger.Error("Server setup failed", "error", err)
		log.Fatal("サーバーの設定に失敗しました:", err)
	}

	if err := internal.StartServerWithGracefulShutdown(server); err != nil {
		logger.Error("Server failed", "error", err)
	}
}
