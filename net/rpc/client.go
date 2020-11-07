package rpc

import (
	"context"

	"google.golang.org/grpc"
)

// Dial connects to addr with opts.
func Dial(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, addr, opts...)
}
