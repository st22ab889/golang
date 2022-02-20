/*
表示符概念
	1. Golang对各种变量、方法等命名时使用的字符序列称为标识符
	2. 凡是自己可以起名字的地方都叫标识符

标识符的命名规则
	1. 由26个英文字母大小写、0-9、_ 组成
	2. 数字不可以开头, 严格区分大小写, 标识符不能包含空格
	3. 下划线"_"本身在Golang中是一个特殊的标识符,称为空标识符.可以代表任何其它的标识符,但是它对应的值会被忽略(比如忽略某个返回值).所以仅能被作为占位符使用,不能作为标识符使用.
	4. 不能以系统保留关键字作为标识符,比如 break、if等等

标识符的命名规范
	1. 包名:保持 package 的名字和目录名称一致,尽量采取有意义、简短的包名,不要和标准库冲突
	2. 变量名、函数名、常量名: 采用驼峰法
	3. (重要)如果变量名、函数名、常量名首字母大写,则可以被其它的包访问.如果首字母小写则只能在本包中使用(注:可以简单的理解成:首字母大写是公有的,首字母小写是私有的)
*/

/*
系统保留关键字:在Golang中为了简化代码编译过程中对代码的解析,其定义的保留关键字只有25个.详见如下:
	break		default			func		interface		select
	case		defer			go			map				struct
	chan		else			goto		package			switch
	const		fallthrough		if			range			type
	continue	for				import		return			var

预定标识符:除了保留关键字外,Golang还提供了36个预定的标识符,其中包括基础数据类型和系统内嵌函数
	append		bool		byte		cap			close		complex
	complex64	complex128	uint16		copy		false		float32
	float64		imag		int			int8		int16		uint32
	int32		int64		iota		len			make		new
	nil			panic		uint64		print		println		real
	recover		string		true		uint		uint8		uintprt
*/

package main
import(
	"fmt"
	_ "unsafe"
	_ "strconv"

	// 包名是从 $GOPATH/src/ 后开始计算的, 使用/进行路径分隔
	"go_code/3-var/demo-09-identifier/utils"
)

func main()  {

	var num int = 88
	var Num int = 99
	fmt.Printf("num=%v, Num=%v \n", num, Num)	// num=88, Num=99
	
	// _ 是空标志符,用于占位
	var _ int = 40
	// fmt.Printf("_ = %v", _)	// 这里就不能这样使用

	var int int = 99	// 使用Golang保留的关键字作为变量名称,语法可以通过,但是工作中要求不能这么做
	fmt.Printf("int=%v \n", int)	// int=99

	// 使用 utils 包中的 add 方法
	fmt.Printf("9 + 90 = %v \n", utils.Add(9,90))	// 9 + 90 = 99

	// 变量在同一个作用域(同一个函数或者代码块)内不能重名
}


/*
1. windows上 go run *.go报错？
	windows的cmd等不支持*.go这种表述, 可以下载 git bash 运行
	windows下 go run ./ 等同于 go run *.go

2. 查看Go环境命令: go env

3. GO语言：同个包下不同文件不能互相调用函数: https://www.freesion.com/article/42441148420/

4. Golang学习之同一个package中函数互相调用的问题: https://blog.csdn.net/yxys01/article/details/77834211/

5. GoLand中报错package xxx is not in GOROOT: https://blog.csdn.net/THEGREATHXY/article/details/109337283
	 gomod和gopath是两个包管理方案,并且相互不兼容,在gopath查找包,按照goroot和多gopath目录下 src/xxx 依次查找
	 在gomod下查找包,解析go.mod文件查找包,mod包名就是包的前缀,里面的目录就后续路径了.在gomod模式下,查找包就不会去gopath查找,只是gomod包缓存在gopath/pkg/mod里面.
	 导入gopath目录下的包时报错"package xxx is not in GOROOT",编译器没有去gopath下找包,原因是GO111MODULE有关,解决方案如下:
	 	go env -w GO111MODULE=off

6. go run命令后面所有*.go文件都必须在项目的根目录下，不能有子目录: https://blog.csdn.net/pengpengzhou/article/details/106620271
*/
