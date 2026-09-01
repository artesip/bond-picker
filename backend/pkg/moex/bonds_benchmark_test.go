package moex

import (
	"context"
	"testing"
)

func BenchmarkGetBonds(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := GetBonds(ctx)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
