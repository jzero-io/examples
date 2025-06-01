package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	ValidateAll() error
}

func ValidatorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if r, ok := req.(Validator); ok {
		if err := r.ValidateAll(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return handler(ctx, req)
}
