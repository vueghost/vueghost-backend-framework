package Security

import "testing"

func TestSecurity_SignKey(t *testing.T) {
	if SignKey == "" {
		t.Fatal("Sign key must not be empty")
	}
}
