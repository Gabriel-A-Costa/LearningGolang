package middleware

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	res, err := handler(ctx, req)

	log.Printf(
		"Method: %s | Time: %s | Error: %v",
		info.FullMethod,
		time.Since(start),
		status.Convert(err).Message(),
	)
	return res, err
}
