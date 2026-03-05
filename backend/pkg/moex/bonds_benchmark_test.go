package moex

import "testing"

func BenchmarkGetBonds(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := GetBonds()
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
