// Package handler contain OpenAPI handlers for Fibonacci API.
package handler

import (
	"context"
	"proGO/api"
	"proGO/internal/fib"
	"proGO/internal/storage"
)

type FibHandler struct {
	store storage.Storer
}

func NewFibHandler(s storage.Storer) *FibHandler {
	return &FibHandler{store: s}
}

func (h *FibHandler) GetFib(ctx context.Context, request api.GetFibRequestObject) (api.GetFibResponseObject, error) {
	n := request.Params.N
	if n < 0 {
		errMsg := "n должно быть ≥ 0"
		return api.GetFib400JSONResponse{
			Error: &errMsg,
		}, nil
	}

	if val, err := h.store.Get(ctx, n); err != nil {
		return nil, err
	} else if val != nil {
		s := val.String()
		return api.GetFib200JSONResponse{
			Result: &s,
		}, nil
	}

	result, err := fib.Fibonacci(n)
	if err != nil {
		msg := err.Error()
		return api.GetFib400JSONResponse{Error: &msg}, nil
	}

	if err := h.store.Set(ctx, n, result); err != nil {
		return nil, err
	}

	s := result.String()
	return api.GetFib200JSONResponse{
		Result: &s,
	}, nil
}
