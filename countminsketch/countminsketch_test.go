package countminsketch_test

import (
	"perhaps/countminsketch"
	"testing"
)

func TestUpdate(t *testing.T) {
	type element struct {
		key string
		val uint64
	}

	for _, tt := range []struct {
		els   []element
		exact bool
	}{
		{exact: true, els: []element{{"a", 1}, {"b", 1}}},
		{exact: false, els: []element{{"a", 1}, {"b", 1}, {"c", 1}}},
	} {
		countMin := countminsketch.New(5, 2)

		actual := map[string]uint64{}

		for _, el := range tt.els {
			actual[el.key] += el.val
			countMin.Add([]byte(el.key), el.val)
		}

		for k, v := range actual {
			est := countMin.Estimate([]byte(k))

			if tt.exact {
				if est != v {
					t.Errorf("estimate %v incorrect wanted %v, got %v", k, v, est)
				}
			} else {
				if est < v {
					t.Errorf("estimate %v to low wanted %v, got %v", k, v, est)
				}
			}
		}
	}
}
