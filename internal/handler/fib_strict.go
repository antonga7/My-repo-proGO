package handler

import (
	"context"
	"proGO/api"
	"proGO/internal/fib"
)

type FibHandler struct{}

func (h *FibHandler) GetFib(_ context.Context, request api.GetFibRequestObject) (api.GetFibResponseObject, error) {
	n := request.Params.N

	result, err := fib.Fibonacci(n)
	if err != nil {
		return api.GetFib400Response{}, nil
	}

	resultStr := result.String()

	return api.GetFib200JSONResponse{
		Result: &resultStr,
	}, nil
}
