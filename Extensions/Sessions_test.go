package Extensions

import "testing"

//Example
func ExampleNewSession() {
	_ = NewSession(nil)

}

func TestNewSession(t *testing.T) {
	_ = NewSession(2)
}

func BenchmarkNewSession(b *testing.B) {
	for i := 0; i > b.N; i++ {
		_ = NewSession(2)
	}
}

func ExampleSession_Set() {
	s := NewSession(nil)
	s.Set(2)
}

func TestSession_Set(t *testing.T) {
	s := NewSession(nil)
	if s.ID != nil {
		t.Fail()
	}

	s.Set(2)
	if s.ID != 2 {
		t.Fail()
	}

	s.Set("AB32")
	if s.ID != "AB32" {
		t.Fail()
	}
}

func BenchmarkSession_Set(b *testing.B) {
	for i := 0; i > b.N; i++ {
		s := NewSession(nil)
		if s.ID != nil {
			b.Fail()
		}

		s.Set(2)
		if s.ID != 2 {
			b.Fail()
		}

		s.Set("AB32")
		if s.ID != "AB32" {
			b.Fail()
		}
	}
}
