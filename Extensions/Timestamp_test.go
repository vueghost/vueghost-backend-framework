package Extensions

import "testing"

func ExampleNewTimestamp() {
	_ = NewTimestamp()
}
func ExampleTimestamp_Now() {
	// Here some example about this function.
	NewTimestamp().Now("2006-02-01")

	//@Returns.
}

func TestTimestamp_Now(t *testing.T) {
	t.Run("default timestamp now", func(t *testing.T) {
		ts := NewTimestamp()
		t.Log(ts.Now())
	})

	t.Run("with timestamp format", func(t *testing.T) {
		ts := NewTimestamp()
		t.Log(ts.Now("2006-02-01"))
	})
}

func BenchmarkTimestamp_Now(b *testing.B) {
	b.Run("default timestamp now", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ts := NewTimestamp()
			b.Log(ts.Now())
		}

	})

	b.Run("with timestamp format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ts := NewTimestamp()
			b.Log(ts.Now("2006-02-01"))
		}
	})
}
