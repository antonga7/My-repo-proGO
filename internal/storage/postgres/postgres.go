// Package postgres предоставляет реализацию интерфейса storage.Storer для базы данных PostgreSQL.
package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/big"
)

// Store реализует интерфейс Storer, используя базу данных PostgreSQL.
type Store struct {
	pool *pgxpool.Pool
}

// New создает новый экземпляр Store с указанной строкой подключения.
func New(connStr string) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к Postgres: %w", err)
	}
	return &Store{pool: pool}, nil
}

// Get получает значение из БД
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
	if _, ok := val.SetString(valStr, 10); !ok {
		return nil, fmt.Errorf("неправильное bigint в БД")
	}

	return val, nil
}

// Set записывает значение в БД
func (s *Store) Set(ctx context.Context, n int, val *big.Int) error {
	_, err := s.pool.Exec(ctx, "INSERT INTO fibonacci(n, value) VALUES ($1, $2) ON CONFLICT (n) DO NOTHING", n, val.String())
	if err != nil {
		return fmt.Errorf("ошибка при сохранении: %w", err)
	}
	return nil
}

func (s *Store) Pool() *pgxpool.Pool {
	return s.pool
}
