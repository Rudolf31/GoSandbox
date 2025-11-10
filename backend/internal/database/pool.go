package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"interface_lesson/internal/config"
)

func NewPool(lc fx.Lifecycle, cfg *config.Config) *pgxpool.Pool {
	pgxCfg, err := pgxpool.ParseConfig(cfg.DB.URL)
	if err != nil {
		panic(err.Error())
	}

	pgxCfg.MinConns = cfg.DB.MinConns
	pgxCfg.MaxConns = cfg.DB.MaxConns

	pool, errPool := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if errPool != nil {
		panic(errPool.Error())
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		panic(err.Error())
	}

	lc.Append(
		fx.StopHook(func() {
			pool.Close()
		}))

	return pool
}
