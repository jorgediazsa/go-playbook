package testingadv

import "testing"

func TestCountCommas(t *testing.T) {
	if CountCommas("a,b,c") != 2 {
		t.Fatal("Logic broken")
	}
	if CountCommas("nolines") != 0 {
		t.Fatal("Logic broken")
	}
	if CountCommas(",,,,") != 4 {
		t.Fatal("Logic broken")
	}
}

// Run with: go test -bench=BenchmarkCountCommas -benchmem
func BenchmarkCountCommas(b *testing.B) {
	s := "word1,word2,some long phrase,word4,word5,word6,word7,word8"

	// Reset timer ensures that setup time (creating the string) isn't counted
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = CountCommas(s)
	}
}
