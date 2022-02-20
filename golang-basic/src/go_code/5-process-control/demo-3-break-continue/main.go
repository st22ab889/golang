
/*
跳转控制语句-break: 
	break语句用于终止某个语句块的执行,用于中断当前for循环或跳出switch语句.
	break语句出现在多层嵌套的语句块中时,可以通过标签指明要终止的是哪一层语句块

跳转控制语句-continue: 
	continue语句用于结束本次循环,继续执行下一次循环
	continue语句出现在多层嵌套的语句块中时,可以通过标签指明要跳过的是哪一层循环

跳转控制语句-goto: 在Golang程序设计中一般不主张使用goto语句
	Golang语言的 goto 语句可以无条件地转移到程序中指定的行
	goto语句通常与条件语句配合使用.可以用来实现条件转移,跳出循环体等功能
	在Golang程序设计中一般不主张使用goto语句,以免造成程序流程的混乱,使理解和调试程序都产生困难

跳转控制语句-return:return使用在方法,表示跳出所在的函数或方法
	如果return是在普通的函数,则表示跳出函数,即不再执行函数中 return 后面的代码,也可以理解为终止函数
	如果return是在main函数,表示终止main函数,也就是说终止程序
*/

package main
import(
	"fmt"
	"math/rand"
	"time"
	_ "unsafe"
	_ "strconv"
)

func main() {
	
	// testBreak()

	// testContinue()
	
	testGoto()
}

func testBreak() {
	
	// 随机生成 1-100 的一个数,直到生成99这个数,看看一共用了几次.该需求可以说明其它流程控制数据的必要性, 比如break
	var count int = 0
	for{
		// 在Golang中为了生成一个随机数,还需要为rand设置一个种子,否则返回的值总是固定的
		// time.Now().Unix() 返回一个从 1970:01:01 的0时0分0秒到现在的秒数
		rand.Seed(time.Now().UnixNano())
		// 随机生成 1-100 之间的整数
		n := rand.Intn(100) + 1		// [0, 100)
		fmt.Println("n=", n)
		count++
		if(n == 99){
			break	// 表示跳出for循环
		}
	}
	fmt.Println("生成99一共使用了 ", count)

	// break语句出现在多层嵌套的语句块中时,可以通过标签指明要终止的是哪一层语句块
	flag1:	// 设置一个标签,这个标签可以自定义名称
	for i:= 0; i < 3; i++{
		flag2:
		for j:= 0; j < 3; j++{
			if j == 2 {
				// break		// break默认跳出最近的for循环
				break flag1 	// break后面可以指定标签,跳出标签对应的for循环
			}else if j == 3 {
				break flag2	
			}
			fmt.Println("i=", i, ",j=", j)
		}
	}
}

func testContinue(){

	for i:= 0; i < 3; i++{
		for j:= 0; j < 3; j++{
			if j == 2 {
				continue
			}
			fmt.Println("i=", i, ",j=", j)
		}
	}
}

func testGoto(){

	var n int = 30
	fmt.Println("ok-1")	// ok-1
	if n > 20{
		goto label_1
	}
	fmt.Println("ok-2")
	label_1:
	fmt.Println("ok-3")	// ok-3

}