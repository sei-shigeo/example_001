# 変数定義
GO_CMD := go
TEMPL_CMD := go tool templ
GOOSE_CMD := go tool goose
SQLC_CMD := go tool sqlc
AIR_CMD := go tool air

# ポート設定
APP_PORT := 8080

# ディレクトリ設定
CMD_SITE_DIR := ./cmd/web-site
CMD_APP_DIR := ./cmd/web-app
CSS_INPUT := ./assets/css/input.css
CSS_OUTPUT := ./assets/css/output.css
TS_INPUT := ./assets/ts/index.ts
JS_OUTPUT := ./assets/js/

# Air設定用変数
AIR_BUILD_DELAY := 100
AIR_EXCLUDE_DIR := node_modules
AIR_INCLUDE_EXT := go
AIR_STOP_ON_ERROR := false
AIR_CLEAN_ON_EXIT := true

# 環境設定
ENV ?= development
ifeq ($(ENV),production)
	BUILD_FLAGS := -ldflags="-s -w"
else
	BUILD_FLAGS :=
endif

# デフォルトターゲット
.DEFAULT_GOAL := help

.PHONY: help
help: ## このヘルプを表示
	@echo "利用可能なコマンド:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# 開発サーバー
templ/app: ## アプリサーバーを起動
	@$(TEMPL_CMD) generate --watch --proxy="http://localhost:$(APP_PORT)" --cmd="go run $(CMD_APP_DIR)/." --open-browser=false

tailwind: ## TailwindCSSを実行
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch --minify

dev: ## アプリの開発環境を起動（templ + server + tailwind）
	make -j2 templ/app tailwind

# データベース管理
mig/create: ## 新しいマイグレーションを作成
	@read -p "Migration name: " name; \
	$(GOOSE_CMD) -s create $$name sql

mig/up: ## マイグレーションを実行
	$(GOOSE_CMD) up

mig/down: ## マイグレーションをロールバック
	$(GOOSE_CMD) down

mig/status: ## マイグレーションの状態を確認
	$(GOOSE_CMD) status

mig/reset: ## マイグレーションをリセット
	$(GOOSE_CMD) reset
	$(GOOSE_CMD) up

# SQLC
sqc/gen: ## SQLCでコードを生成
	$(SQLC_CMD) generate

sqc/comp: ## SQLCでコンパイル
	$(SQLC_CMD) compile

# ビルド
build/site: sqc/gen ## サイト用バイナリをビルド
	$(GO_CMD) build $(BUILD_FLAGS) -o bin/site $(CMD_SITE_DIR)

build/app: sqc/gen ## アプリ用バイナリをビルド
	$(GO_CMD) build $(BUILD_FLAGS) -o bin/app $(CMD_APP_DIR)

build: build/site build/app ## 全てのバイナリをビルド

# テスト
test: ## テストを実行
	$(GO_CMD) test -v ./...

test/coverage: ## カバレッジ付きでテストを実行
	$(GO_CMD) test -v -coverprofile=coverage.out ./...
	$(GO_CMD) tool cover -html=coverage.out -o coverage.html

# クリーンアップ
clean: ## 生成されたファイルを削除
	rm -f $(CSS_OUTPUT)
	find . -name "*.templ.go" -delete

clean/all: clean ## 全ての生成ファイルとキャッシュを削除
	$(GO_CMD) clean -cache
	$(GO_CMD) clean -modcache
	rm -rf bin/
	rm -f coverage.out coverage.html