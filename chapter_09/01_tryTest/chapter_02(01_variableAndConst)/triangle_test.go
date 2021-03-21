package main

import (
	"testing"
	"time"
)

func TestTriangle(t *testing.T) {
	// 定义测试数据
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	time.Sleep(time.Microsecond)

	// 遍历测试数据，测试calcTriangle
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

/* 测试通过终端输出
➜  chapter_02(01_variableAndConst) git:(master) ✗ go test .
ok      testTriangle    0.293s
*/

/*  错误时终端的输出
➜  chapter_02(01_variableAndConst) git:(master) ✗ go test .
--- FAIL: TestTriangle (0.00s)
    triangle_test.go:16: calcTriangle(3, 4); got 5; expected 6
    triangle_test.go:16: calcTriangle(8, 15); got 17; expected 16
FAIL
FAIL    testTriangle    0.285s
FAIL
*/
