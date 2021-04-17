package bloom_test

import (
	"fmt"
	"perhaps/bloom"
	"testing"
)

func TestAdd(t *testing.T) {
	bf := bloom.New(3, 2)
	bf.Add([]byte("foo"))

	fmt.Println("checking foo", bf.Check([]byte("foo")))
	fmt.Println("checking bar", bf.Check([]byte("bar")))

	bf.Add([]byte("baz"))

	fmt.Println("checking foo", bf.Check([]byte("foo")))
	fmt.Println("checking bar", bf.Check([]byte("bar")))

	bf.Add([]byte("buzz"))

	fmt.Println("checking foo", bf.Check([]byte("foo")))
	fmt.Println("checking bar", bf.Check([]byte("bar")))

	for _, tt := range []struct {
		vals  []string
		check string
		want  bool
	}{
		{vals: []string{"foo"}, check: "foo", want: true},
		{vals: []string{"a"}, check: "not-foo", want: false},
		{vals: []string{"a", "b", "c"}, check: "not-foo", want: false},
		{vals: []string{"a", "b", "c", "d"}, check: "not-foo", want: false},
		{vals: []string{"a", "b", "c", "d", "e"}, check: "not-foo", want: true},
	} {
		bf := bloom.New(6, 2)

		for _, val := range tt.vals {
			bf.Add([]byte(val))
		}

		if got := bf.Check([]byte(tt.check)); got != tt.want {
			t.Errorf("checked %v, wanted %v got %v", tt.check, tt.want, got)
		}
	}
}
