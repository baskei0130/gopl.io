package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("ab"), []byte("ba")},
		{[]byte("abc"), []byte("cba")},
		{[]byte("阿吽"), []byte("吽阿")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s", string(test.b))
		reverse(test.b)
		if string(test.b) != string(test.want) {
			t.Errorf("case %q: got %q, want %q", descr, string(test.b), string(test.want))
		}
	}
}
