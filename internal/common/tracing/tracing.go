package tracing

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"google.golang.org/grpc/metadata"
)

func generateTracingID() string {
	size := 16 // 128 bits
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func NewContextWithTracingID(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["request_id"]) == 0 {
		newMD := metadata.Join(md, metadata.Pairs("request_id", generateTracingID()))
		return metadata.NewIncomingContext(context.Background(), newMD)
	}
	return ctx
}

func NewOutgoingContextWithTracingID(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["request_id"]) == 0 {
		return ctx
	}
	return metadata.NewOutgoingContext(ctx, md)
}
