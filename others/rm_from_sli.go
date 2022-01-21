package main

import (
	"reflect"
	"testing"
)

const (
	right = "\u2713"
	wrong = "\u2715"
)

func removeItem(sli []int, n int) []int {
	for i, tn := range sli {
		if tn == n {
			sli = append(sli[:i], sli[i+1:]...)
		}
	}
	return sli
}

func TestRemoveItem(t *testing.T) {
	testCases := []struct {
		name   string
		inpSli []int
		inpN   int
		expect []int
	}{
		{
			name:   "not_exist",
			inpSli: []int{1, 3, 4},
			inpN:   2,
			expect: []int{1, 3, 4},
		},
		{
			name:   "exist",
			inpSli: []int{1, 2, 3},
			inpN:   2,
			expect: []int{1, 3},
		},
		{
			name:   "remove_the_last_one",
			inpSli: []int{2},
			inpN:   2,
			expect: []int{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("\tTestCase: %s", tc.name)
			{
				get := removeItem(tc.inpSli, tc.inpN)
				if !reflect.DeepEqual(get, tc.expect) {
					t.Fatalf("\t%s\tcase failed: got(%v) != expect(%v)",
						wrong, get, tc.expect)
				}
				t.Logf("\t%s\tcase passed: got(%v) == expect(%v)",
					right, get, tc.expect)
			}
		})
	}
}
