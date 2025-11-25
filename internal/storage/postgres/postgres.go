package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/big"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(connStr string) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к Postgres: %w", err)
	}
	return &Store{pool: pool}, nil
}

func (s *Store) Get(ctx context.Context, n int) (*big.Int, error) {
	var valStr string
	err := s.pool.QueryRow(ctx, "SELECT value FROM fibonacci WHERE n = $1", n).Scan(&valStr)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	val := new(big.Int)
	val.SetString(valStr, 10)
	return val, nil
}

func (s *Store) Set(ctx context.Context, n int, val *big.Int) error {
	_, err := s.pool.Exec(ctx, "INSERT INTO fibonacci(n, value) VALUES ($1, $2) ON CONFLICT (n) DO NOTHING", n, val.String())
	if err != nil {
		return fmt.Errorf("ошибка при сохранении: %w", err)
	}
	return nil
}
