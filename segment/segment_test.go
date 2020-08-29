package segment

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	nums := []interface{}{-2, 0, 3, -5, 2, -1}
	test := [][]interface{}{
		{0, 2, 1},
		{2, 4, 0},
		{3, 5, -4},
	}

	s := New(nums, func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	for _, v := range test {
		if s.Query(v[0].(int), v[1].(int)) != v[2] {
			t.Error("线段树查询错误")
		}
	}

	s.Set(0, -1)
	s.Set(2, 5)
	s.Set(5, -3)

	test = [][]interface{}{
		{0, 2, 4},
		{2, 4, 2},
		{3, 5, -6},
	}
	for _, v := range test {
		if s.Query(v[0].(int), v[1].(int)) != v[2] {
			t.Error("线段树修改错误")
		}
	}
}