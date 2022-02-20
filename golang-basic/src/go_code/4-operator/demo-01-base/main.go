
/*
1) 算术运算符: 对数值类型的变量进行运算,比如加减乘除
	+(正号)、-(负号)、+(加)、-(减)、*、/、%(取模/余)、++、--、+(字符串相加)
2) 赋值运算符:赋值运算符就是将某个运算后的值,赋给指定的变量
	=、+=、-=、*=、/=、%=
	<<= (左移后赋值)、>>= (右移后赋值)、&= (按位与后赋值)、^= (按位异或后赋值)、|= (按位或后赋值)
3) 比较/关系运算符: 关系运算符的结果都是bool型,结果只能是true或false.关系表达式经常用在if结构的条件中或循环结构的条件中
	==、!=、<、>、<=、>=
4) 逻辑运算符: 用于连接多个条件(一般来讲就是关系表达式),最终的结果也是一个bool值
	&&、||、!
5) 位运算符
	& 按位与运算符,"&"是双目运算符.其功能是参与运算的两数各对应的二进位相与.运算规则是:同时为1,结果为1,否则为0
	| 按位或运算符,"|"是双目运算符.其功能是参与运算的两数各对应的二进位相或.运算规则是:有一个为1,结果为1,否则为0
	^ 按位异或运算符,"^"是双目运算符.其功能是参与运算的两数各对应的二进位相异或.运算规则是:当二进位不同时,结果为1,否则为0
	<< 左移运算符,"<<"是双目运算符.其功能是把"<<"左边的运算数的各二进位全部左移若干位,高位丢弃,低位补0.左移n位就是乘以2的n次方
	>> 右移运算符,">>"是双目运算符.其功能是把">>"左边的运算数的各二进位全部右移若干位.右移n位就是除以2的n次方
6) 其它术运算符
	& 返回变量存储地址
	* 指针变量
7) Golang明确不支持三元运算符.Golang的设计理念是"一种事情有且只有一种方法完成"
*/

package main
import(
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main() {

	// 除后保留整数部分,去掉小数部分
	fmt.Printf("n = %v \n", 10 / 4)		// n = 2 

	var n1 float32  = 10 / 4
	fmt.Printf("n1 = %v \n", n1)		// n1 = 2 

	// 如果希望保留小数部分,则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Printf("n2 = %v \n", n2)		// n2 = 2.5


	// 公式: a % b = a - a / b * b
	fmt.Println("10%3=", 10 % 3)		// 1
	fmt.Println("-10%3=", -10 % 3)		// -1  满足公式: -10 % 3 = -10 - (-10) / 3 * 3 = -10 - (-9) = -1
	fmt.Println("10%-3=", 10 % -3)		// 1
	fmt.Println("-10%-3=", -10 % -3)	// -1

	// ++ 、--
	var i int = 10
	i++ // 等价 i = i + 1
	fmt.Println("i=", i)	// i= 11
	i--
	fmt.Println("i=", i)	// i= 10


	/*
	算术运算符细节说明
	1. 对于"/",它的整数部分和小数部分是有区别的.整数之间做除法时,只保留整数部分而舍弃小数部分.例如 x:=19/5 结果是 3
	2. 当对一个属取模时,可以等价 a % b = a - a / b * b, 这样可以看到取模的一个本质运算
	3. Golang 的自增自减只能当作一个独立语言使用, 不能这样使用 b:=a++ 或者 b:=a--
	4. Golang 的 ++ 和 -- 只能写在变量的后面, 不能写在变量的前面. 只用 a++ 或 a--, 不能 ++a 或 --a
	5. Golang 的设计者去掉 c/java 中的自增自减的容易混淆的写法,让Golang更加简洁、统一。
	*/

	relationalOperator() 

	logicalOperators()

	assignmentOperator() 

	otherOperator()

	inputFromTerminal() 
}

func relationalOperator() {

	var n1 int = 8
	var n2 int = 9
	fmt.Println(n1 == n2)	// false
	fmt.Println(n1 != n2)	// true
	fmt.Println(n1 > n2)	// false
	fmt.Println(n1 >= n2)	// false
	fmt.Println(n1 < n2)	// true
	fmt.Println(n1 <= n2)	// true

	flag := n1 > n2
	fmt.Println("flag = ", flag) // flag =  false

	// 关系运算符组成的表达式称为 关系表达式, 例如 a > b
}


func logicalOperators(){

	var age int = 40
	if age > 30 && age < 50 {	// true
		fmt.Println("abc")
	} 
	
	if age > 30 || age < 40 {	// true
		fmt.Println("abcd")
	} 
	
	if !(age > 30) {			// false
		fmt.Println("abcde")
	}
	
	/*
	逻辑运算符细节说明:
	&& 也叫短路与,如果第一个条件为false,则第二个条件不会判断,最终结果为false
	|| 也叫短路或,如果第一个条件为true,则第二个条件不会判断,最终结果为true
	*/
}

func assignmentOperator(){

	// 复合赋值的操作
	a := 8
	a += 8 		// 等价 a = a + 8
	fmt.Println("a = ", a)	// 	a =  16

	var c int 
	c = a + 3	// 赋值运算符的执行顺序是从右向左
	fmt.Println("c = ", c)	// 	c =  19


	// 赋值运算符的左边只能是变量,右边可以是变量、表达式(任何有值都可以看作是表达式)、常量值
	var d int
	d = a
	d = 8 + 8 * 9	 // = 的右边是表达式
	d = test() + 80	 // = 的右边是表达式
	d = 99		// 99 是常量值
	fmt.Println("d = ", d)	// 	d =  99
}

func test() int {
	return 8
}

func otherOperator() {
	a := 100
	fmt.Println("a的地址 = ", &a)			// a的地址 =  0xc0000aa190
	var ptr *int = &a	
	fmt.Println("ptr指向的值是 = ", *ptr)	// ptr指向的值是 =  100
}

func inputFromTerminal() {
	
	// 获取用户终端输入, 调用 fmt 包的 fmt.Scanln() 或者 fmt.Scanf()

	var name string
	var age byte
	var sal float32
	var isPass bool
	
	// 使用 fmt.Scanln() 从控制台接收用户信息, 当程序执行到 fmt.Scanln() 会停止在这里,等待用户输入,并回车
	/*
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v, 年龄是 %v, 薪水是 %v, 是否通过考试 %v \n", name, age, sal, isPass)
	*/

	// 使用 fmt.Scanf(), 按指定格式输入
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试,使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v, 年龄是 %v, 薪水是 %v, 是否通过考试 %v \n", name, age, sal, isPass)

}