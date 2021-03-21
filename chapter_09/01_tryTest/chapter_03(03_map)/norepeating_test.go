package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s         string
		subStr    string
		subStrLen int
	}{
		// Normal cases
		{"abcabcbb", "abc", 3},
		{"asdgetdtdd", "asdget", 6},

		// edge cases
		{"", "", 0},
		{"b", "b", 1},
		{"bbbbbbb", "b", 1},
		{"abcabcadcd", "bcad", 4},

		// Chinese support
		{"一二三二一", "一二三", 3},
		{"What_are_you弄啥了！What_are_you弄啥了！", "are_you弄啥了！Wh", 14},
	}

	for _, tt := range tests {
		actualSubstr, actualLen := maxNonRepeatSubStr(tt.s)
		if actualSubstr != tt.subStr || actualLen != tt.subStrLen {
			t.Errorf("maxNonRepeatSubStr: %s, got actual Substr: %s, got actual subStrLen: %d \n"+
				"expected Substr: %s, expected subStrLen: %d", tt.s, actualSubstr, actualLen, tt.subStr, tt.subStrLen)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "What_are_you弄啥了！What_are_you弄啥了！"
	subStr := "are_you弄啥了！Wh"
	subStrLen := 14
	for i := 0; i< b.N; i++{
		actualSubstr, actualLen := maxNonRepeatSubStr(s)
		if actualSubstr != subStr || actualLen != subStrLen {
			b.Errorf("maxNonRepeatSubStr: %s, got actual Substr: %s, got actual subStrLen: %d \n"+
				"expected Substr: %s, expected subStrLen: %d", s, actualSubstr, actualLen, subStr, subStrLen)
		}
	}
}

/* 运行Benchmark查看运行性能 命令：go test -bench .
goos: darwin
goarch: amd64
pkg: testNoRepeating
BenchmarkSubstr-12        374696              3075 ns/op
PASS
ok      testNoRepeating 1.552s
*/