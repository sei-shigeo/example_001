package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config アプリケーションの設定構造体
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	App      AppConfig      `mapstructure:"app"`
	Log      LogConfig      `mapstructure:"log"`
}

// ServerConfig サーバー設定
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseConfig データベース設定
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

// AppConfig アプリケーション設定
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

var GlobalConfig *Config

// LoadConfig 設定ファイルを読み込む
func LoadConfig(configPath string) (*Config, error) {
	// 1. .envファイルを読み込んで環境変数に設定
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// 2. 設定ファイルの読み込み
	if configPath != "" {
		// 完全なファイルパスが指定された場合
		viper.SetConfigFile(configPath)
	} else {
		// デフォルトのパスにあるconfig.tomlファイルを読み込み
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath("./internal/config")
		viper.AddConfigPath(".")
	}

	// 3. 環境変数を読み込み
	viper.AutomaticEnv()

	// 4. 環境変数のマッピングを設定
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.user", "DATABASE_USER")

	// 5. デフォルト値を設定
	setDefaults()

	// 6. config.tomlを読み込み
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("設定ファイルが見つかりません。デフォルト値を使用します: %v", err)
		} else {
			return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
		}
	} else {
		log.Printf("設定ファイルを読み込みました: %s", viper.ConfigFileUsed())
	}

	// 7. 構造体にマッピング
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("設定の解析に失敗しました: %w", err)
	}

	// 8. グローバル設定を更新
	GlobalConfig = &config
	return &config, nil
}

// setDefaults デフォルト値を設定
func setDefaults() {
	// サーバー設定のデフォルト値
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)

	// データベース設定のデフォルト値
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.dbname", "myapp")
	viper.SetDefault("database.sslmode", "disable")

	// アプリケーション設定のデフォルト値
	viper.SetDefault("app.name", "My Go App")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("app.debug", true)

	// ログ設定のデフォルト値
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "logs/app.log")
}

// GetConfig グローバル設定を取得
func GetConfig() *Config {
	return GlobalConfig
}

// GetServerAddr サーバーアドレスを取得
func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// GetDatabaseDSN データベース接続文字列を取得
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// Validate 設定を検証する
func (c *Config) Validate() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", c.Database.Port)
	}

	if c.App.Environment != "development" &&
		c.App.Environment != "staging" &&
		c.App.Environment != "production" {
		return fmt.Errorf("invalid environment: %s", c.App.Environment)
	}

	return nil
}
