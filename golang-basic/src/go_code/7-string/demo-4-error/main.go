/*
Golang的错误处理机制:
	1.默认情况下,当发生错误后程序就会退出(崩溃).
	2.当发生错误后也可以捕获到错误并进行处理,保证程序可以继续执行.还可以在捕获错误后,给管理员一个提示(邮件、短信 ...)
	3.Golang不支持传统的 try...catch...finally 这种处理
	4.Golang中引入的处理方式为: defer, panic, recover
	5.这几个异常的使用场景可以这么简单描述:Golang中可以抛出一个panic的异常,然后在defer中通过recover捕获这个异常,然后正常处理

错误处理的好处:
	进行错误处理后,程序不会轻易的刮掉,如果假如预警代码,就可以让程序更加的健壮

自定义错误:Golang中也支持自定义错误,使用 errors.New 和 panic 内置函数
	errors.New("错误说明"),会返回一个error类型的值,表示一个错误
	panic内置函数接收一个interface{}类型的值(也就说任何值)作为参数.可以接收error类型的变量,输出错误信息,并退出程序
*/


package main
import(
	"fmt"
	"errors"
	_ "strconv"
	_ "strings"
	_ "time"
)

func test(){
	defer func(){
		err := recover()	// recover()是内置函数,可以捕获到异常
		if err != nil {	// 说明捕获到错误
			fmt.Println("err=", err)
			fmt.Println("这里就可以将异常发给管理员", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)	
}

// 需求:函数去读取以配置文件 init.conf 的信息,如果文件名传入不正确就返回一个自定义的错误
func readConf(name string) (err error){
	if name == "init.conf" {
		// 读取...
		return nil
	}else{
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test02(){
	err := 	readConf("init2.conf")
	if err != nil {
		// 如果读取文件发生错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")	
}

func main() {

	/*
	// 测试1:捕获错误
	test();	
	for {	// for 循环
		fmt.Println("执行完成....")	
		time.Sleep(time.Second)
	}
	*/

	
	// 测试2:测试自定义错误
	test02()
	fmt.Println("继续下面的代码....")	
	/*
	执行结果如下:
		panic: 读取文件错误...

		goroutine 1 [running]:
		main.test02()
				D:/GoProject/src/go_code/7-string/demo-4-error/main.go:56 +0x49
		main.main()
				D:/GoProject/src/go_code/7-string/demo-4-error/main.go:73 +0x19
		exit status 2
	*/
}