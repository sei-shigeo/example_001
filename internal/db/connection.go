package db

import (
	"context"
	"fmt"
	"time"

	"project/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB データベース接続を管理する構造体
type DB struct {
	Pool    *pgxpool.Pool
	Queries *Queries
}

// NewDB 新しいデータベース接続を作成
func NewDB(config *configs.Config) (*DB, error) {
	// データベース接続文字列を取得
	dsn := config.GetDatabaseDSN()

	// 接続プールの設定
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("データベース設定の解析に失敗しました: %w", err)
	}

	// 接続プールの設定をカスタマイズ
	poolConfig.MaxConns = 10                      // 最大接続数
	poolConfig.MinConns = 2                       // 最小接続数
	poolConfig.MaxConnLifetime = time.Hour        // 接続の最大生存時間
	poolConfig.MaxConnIdleTime = time.Minute * 30 // アイドル接続の最大時間

	// データベース接続プールを作成
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗しました: %w", err)
	}

	// 接続をテスト
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("データベース接続のテストに失敗しました: %w", err)
	}

	// Queriesインスタンスを作成
	queries := New(pool)

	return &DB{
		Pool:    pool,
		Queries: queries,
	}, nil
}

// Close データベース接続を閉じる
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// HealthCheck データベースのヘルスチェック
func (db *DB) HealthCheck(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}

// WithTx トランザクション内でクエリを実行
func (db *DB) WithTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("トランザクションの開始に失敗しました: %w", err)
	}
	defer tx.Rollback(ctx)

	queries := db.Queries.WithTx(tx)
	if err := fn(queries); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("トランザクションのコミットに失敗しました: %w", err)
	}

	return nil
}
