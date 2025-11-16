package storage

import (
	"context"
	"math/big"
)

type Storer interface {
	Get(ctx context.Context, n int) (*big.Int, error)
	Set(ctx context.Context, n int, val *big.Int) error
}
