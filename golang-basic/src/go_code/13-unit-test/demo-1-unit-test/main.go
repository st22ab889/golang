/*
1.单元测试:如何确认一个函数或者一个模块的结果是否正确?
	传统方法:调用函数看实际输出的结果是否和预期的结果一致.但是有以下缺点:
		(1):需要在main函数中调用,这就需要修改main函数,不方便. 如果项目正在运行就可能去停止项目.
		(2):不利于管理,因为当测试多尔函数或者多个模块时,都需要写在main函数.
2.Go语言中自带有一个轻量级的测试框架"testing"和自带的"go test"命令来实现单元测试和性能测试.
  "testing"框架和其它语言中的测试框架类似,可以基于这个框架写针对相应函数的测试用例,也可以基于该框架写相应的压力测试用例.
  通过单元测试可以解决如下问题:
	(1):确保每个函数是可运行,并且运行结果是正确的.
	(2):确保写出来的代码性能是好的.
	(3):单元测试能及时的发现程序设计或实现的逻辑错误,使问题及早暴露,便于问题的定位解决;
	    而性能测试的重点在于发现程序设计上的一些问题,让程序能够在高并发的情况下还能保持稳定;
3.特别说明:测试时,可能需要暂时退出360,因为360可能会认为生成的测试用例程序是木马.

4.单元测试快速入门总结:
	(1):测试用例文件名必须以 _test.go 结尾.比如 cal_test.go, cal 不是固定的
	(2):测试用例函数必须以 Test 开头, 一般来说就是 Test+被测试的函数名, 比如 TestAddUpper
	(3):TestAddUpper(t *testing.T)的形参类型必须是 *testing.T (可以参考手册)
	(4):一个测试用例文件中可以有多个测试用例函数,比如 TestAddUpper、TestSub; 
	(5):运行测试用例指令:
			go test		[如果运行正确,无日志,错误时会输出日志]
			go test	-v	[运行正确或是错误都输出日志]
	(6):当出现错误时,可以使用 t.Fatalf来格式化输出错误信息,并退出程序
	(7):t.Logf方法可以输出相应的日志
	(8):测试用例函数,并没有放在main函数中也执行了,这就是测试用例的方便之处
	(9):PASS表示测试用例运行成功,FAIL表示测试用例运行失败
	(10):测试单个文件,一定要带上被测试的原文件
		go test -v cal_test.go main.go
	(11):测试单个方法(注意:下面这个命令运行出错,暂时没有找到解决方法)
		go test -v -test.run TestAddUpper
	(12):对包进行测试
		go test 包名称
	(13):go test 命令介绍: https://www.jianshu.com/p/9b4af198ee28
	(14):"testing"框架还可以测试代码覆盖率.
	(15): Run 'go help test' and 'go help testflag' for details.

5. type T
	func (c *T) Error(args ...interface{})
	func (c *T) Errorf(format string, args ...interface{})
	func (c *T) Fail()
	func (c *T) FailNow()
	func (c *T) Failed() bool
	func (c *T) Fatal(args ...interface{})
	func (c *T) Fatalf(format string, args ...interface{})
	func (c *T) Log(args ...interface{})
	func (c *T) Logf(format string, args ...interface{})
	func (c *T) Name() string
	func (t *T) Parallel()
	func (t *T) Run(name string, f func(t *T)) bool
	func (c *T) Skip(args ...interface{})
	func (c *T) SkipNow()
	func (c *T) Skipf(format string, args ...interface{})
	func (c *T) Skipped() bool	
*/

package main

import (
	"fmt"
)

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n - 1; i++ {
		res += i 
	}
	return res
}


func getSub(n1 int, n2 int) int {
	return n1 - n2
}



func main() {
	res := addUpper(10)	
	res2 := getSub(10, 3)

	fmt.Printf("addUpper = %v, getSub = %v", res, res2)
}