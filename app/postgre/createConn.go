package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type ConnInfo struct {
	username string
	password string
	host     string
	port     string
	database string
}

func DoWithtries(fn func() error, attemts int, delay time.Duration) (err error) {
	for attemts < 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemts--

			continue
		}
		return nil
	}
	return
}

func newClient(ctx context.Context, maxAttemts int, conn ConnInfo) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprint("postgresql://%s:%s@%s:%s/%s", conn.username, conn.password,
		conn.host, conn.port, conn.database)

	DoWithtries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttemts, 5*time.Second)

	return
}
