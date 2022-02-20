package main

import (
	_ "fmt"
	"testing"	// 引入Go的testing框架包
)

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)	//调用
	if res != 7 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("getSub(10, 3) 执行错误，期望值=%v 实际值=%v\n", 7, res)
	}
	
	//如果正确，输出日志
	t.Logf("getSub(10, 3) 执行正确!")
}