package Extensions

import "testing"

func ExampleNewSanitize() {
	_ = NewSanitize()
}
func TestNewSanitize(t *testing.T) {
	_ = NewSanitize()
}
func BenchmarkNewSanitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewSanitize()
	}
}

func ExampleSanitize_URLQueryEscape() {
	s := NewSanitize()
	s.URLQueryEscape("hello home web")
}
func TestSanitize_URLQueryEscape(t *testing.T) {
	s := NewSanitize()
	r := s.URLQueryEscape("hello home%web")
	if r != "hello+home%25web" {
		t.Fatal()
	}
	t.Log(r)
}
func BenchmarkSanitize_URLQueryEscape(b *testing.B) {
	s := NewSanitize()
	r := s.URLQueryEscape("hello home%web")
	if r != "hello+home%25web" {
		b.Fatal()
	}
	b.Log(r)
}
