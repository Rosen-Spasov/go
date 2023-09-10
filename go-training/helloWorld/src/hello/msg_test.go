package hello

import "testing"

func Test_SayHello(t *testing.T) {
	err := SayHello()
	if err != nil {
		t.Error(err)
	}
}
