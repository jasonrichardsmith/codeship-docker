package simple

import "testing"

func TestReturnString(t *testing.T) {
	s := ReturnString()
	if s != "hello" {
		t.Error("test failed")
	}
}
