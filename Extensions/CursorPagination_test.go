package Extensions

import "testing"

func TestCursorPagination_startCursor(t *testing.T) {
	type test struct {
		page   int
		limit  int
		result int
	}

	testList := map[string]test{
		"tartCursor@page-1": {
			page:   1,
			limit:  20,
			result: 0,
		},
		"tartCursor@page-2": {
			page:   2,
			limit:  20,
			result: 20,
		},
		"tartCursor@page-3": {
			page:   3,
			limit:  20,
			result: 40,
		},
	}

	c := CursorPagination{}
	for name, ts := range testList {
		t.Run(name, func(t *testing.T) {
			r := c.startCursor(ts.page, ts.limit)
			if r != ts.result {
				t.Fail()
			}
		})
	}
}
func TestCursorPagination_endCursor(t *testing.T) {
	type test struct {
		page   int
		limit  int
		result int
	}

	testList := map[string]test{
		"endCursor@page-1": {
			page:   1,
			limit:  20,
			result: 20,
		},
		"endCursor@page-2": {
			page:   2,
			limit:  20,
			result: 40,
		},
		"endCursor@page-3": {
			page:   3,
			limit:  20,
			result: 60,
		},
	}

	c := CursorPagination{}
	for name, ts := range testList {
		t.Run(name, func(t *testing.T) {
			r := c.endCursor(ts.page, ts.limit)
			if r != ts.result {
				t.Fail()
			}
		})
	}
}
func TestCursorPagination_hasNext(t *testing.T) {
	type testParam struct {
		numOfRows int
		limit     int
		result    bool
	}

	testList := map[string]testParam{
		"endCursor@page-1": {
			numOfRows: 10,
			limit:     10,
			result:    true,
		},
		"endCursor@page-2": {
			numOfRows: 30,
			limit:     20,
			result:    true,
		},
		"endCursor@page-3": {
			numOfRows: 3,
			limit:     4,
			result:    false,
		},
	}

	c := CursorPagination{}
	for name, ts := range testList {
		t.Run(name, func(t *testing.T) {
			r := c.hasNext(ts.numOfRows, ts.limit)
			if r != ts.result {
				t.Fail()
			}
		})
	}
}
