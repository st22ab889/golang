package main

/*
测试方法:打开终端,进入此文件所在目录下, 使用"go test -v"命令.
*/

import (
	"fmt"
	"testing"	// 引入Go的testing框架包
)

/*
 1. "TestXxx",其中Xxx可以是任何字母数字字符串(但是第一个字母不能是[a-z]的小写字母),用来识别测试例程.
 2. 在这些函数中,使用 Error、Fail或相关方法来发出失败信号.
 3. 要编写一个新的测试套件,需要创建一个名称以 _test.go 结尾的文件,该文件包含 'testXxx'函数,如上所述。 
   将该文件放在与被测试的包相同的包中。该文件将被排除在正常的程序包之外，但在运行 “go test” 命令时将被包含。
*/

// 编写测试用例去测试 addUpper 函数是否正确
func TestAddUpper(t *testing.T){
	res := addUpper(10)	//调用
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		
		// 调用 Fatalf 相当于在调用 Logf 之后调用 FailNow .
		// Logf 使用与 Printf 相同的格式化语法对它的参数进行格式化,然后将格式化后的文本记录到错误日志里面.如果输入的格式化文本最末尾没有出现新行,那么将一个新行添加到格式化后的文本末尾
		// 将当前的测试函数标识为“失败”,并停止执行该函数.在此之后,测试过程将在下一个测试或者下一个基准测试中继续.
		// FailNow 必须在运行测试函数或者基准测试函数的 goroutine 中调用,而不能在测试期间创建的 goroutine 中调用.调用 FailNow 不会导致其他 goroutine 停止
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用..")
}

