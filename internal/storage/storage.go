// Package storage реализует интерфейсы работы с БД
package storage

import (
	"context"
	"math/big"
)

// Storer реализует интерфейсов Get и Set для хранилища
type Storer interface {
	Get(ctx context.Context, n int) (*big.Int, error)
	Set(ctx context.Context, n int, val *big.Int) error
}
