/*
init函数
	每一个源文件都可以包含一个init函数,会在main函数前被调用,init函数最主要的作用是完成一些初始化的工作
	如果一个文件同时包含变量定义、init、main, 则执行流程是 变量定义 -> init函数 -> main函数
	

*/



package main
import(
	"fmt"
	"go_code/6-function/demo-3-init-func/utils"
)

// 变量定义
var age = test();

func test() int {
	fmt.Println("test()...")
	return 90
}

// init函数,通常在init函数中完成初始化工作
func init(){
	fmt.Println("init()...")
}


func main() {
	fmt.Println("main()... age = ", age)
	fmt.Println("Age = ", utils.Age, ", Name = ", utils.Name);
}

/*
执行结果:
	Util -> init()
	test()...
	init()...
	main()... age =  90
	Age =  100 , Name =  test

结论:可以看到是先执行 其它 源文件的 变量定义、init函数, 再执行 main 源文件中的 变量定义、init函数、main函数

*/