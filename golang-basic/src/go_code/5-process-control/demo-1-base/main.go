/*
流程控制决定程序如何执行,主要有三大流程控制语句:
	1. 顺序控制: 程序从上到下逐行执行,中间没有任何判断和跳转.注意事项:Golang中定义变量时采用合法的前向引用
	2. 分支控制: 让程序有选择执行
		单分支: Golang的if有一个强大的地方就是条件判断语句里面允许声明一个变量,这个变量的作用域只能在该条件逻辑块内,其它地方就不起作用了
			if 条件表达式 {
				// 注意: {} 时必须有的,就算只写一行代码
			}
		双分支: 只会执行其中的一个分支
			if 条件表达式{
				// 执行的代码			
			}else {
				// 执行的代码
			}
		多分支: 多分支只能有一个执行入口, else分支不是必须的
			if 条件表达式{
				// 执行的代码			
			}else if 条件表达式{
				// 执行的代码
			}
			...
			else{
				// 执行的代码
			}
		嵌套分支:嵌套分支不宜过多,建议最多3层
			if 条件表达式{
				if 条件表达式{
					// 执行的代码
				}else{
					// 执行的代码	
				}
			}
		switch分支结构: switch语句用于基于不同条件执行不同动作,每一个case分支都是唯一的,从上到下逐一判断,直到匹配为止.匹配项后面也不需要再加break
			switch 表达式{
				case 表达式1, 表达式2, ... :	// 只要满足启动一个表达式,就会执行这个case下面的语句块
					语句块1
				case 表达式1, 表达式2, ... :
					语句块2
					// 这里不需要写break,因为默认会有,即在默认情况下,当程序执行完case语句块后,就直接退出该switch控制结构
				// 可以有多个case语句
				default:	// 如果值没有和任何的case的表达式匹配成功,则执行default语句块
					语句块					
			}	

	3. 循环控制:让代码循环执行
		for 循环
			for	循环变量初始化;循环条件;循环变量迭代{
				// 循环操作(语句),也叫循环体
			}
		Golang没有 while 和 do...while 语法(比如 java、c等语言就有),可以通过for循环来实现其使用效果

*/

package main
import(
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main() {

	// testIfElse()

    // testSwitch()

	/*
	swich 和 if 的比较:
	1.如果判断的具体数值不多,而且复合整数、浮点数、字符、字符串这几种类型.建议使用switch语句,简洁高效
	2.其它情况:对区间判断和结果位bool类型的判断,使用if,if的使用范围更广
	*/

	// testFor()

	testWhile()
}	


func testIfElse() {

	// 单分支
	/*
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄是", age)
	}
	*/

	// Golang支持在if中直接定义一个变量
	if age := 20; age > 18 {
		fmt.Println("你已经到了法定年龄")
	}

	// 双分支
	if age := 16; age > 18 {
		fmt.Println("你已经到了法定年龄")
	}else{
		fmt.Println("你还没到法定年龄")
	}

	// 多分支
	/*
	var score int
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("买一个大玩具")
	}else if score > 80 && score <= 99 {
		fmt.Println("买一个比较大玩具")
	}else if score > 60 && score <= 80 {
		fmt.Println("买一个小玩具")
	}else{
		fmt.Println("什么都不买")
	}
	*/

	// 多分支使用陷进
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")	// 输出, 多分支只要找到一个入口执行后,整个多分支就结束了. 所以即使 else if 的条件表达式满足也不会执行
	}else if n >6 {
		fmt.Println("ok2")
	}else if n >6 {
		fmt.Println("ok3")
	}else{
		fmt.Println("ok4")
	}
}

func testSwitch() {
	/*
	switch细节讨论:
	1. case后是一个表达式(即:常量值、变量、一个有返回值的函数都可以)
	2. case后的各个表达式的值的数据类型,必须和switch的表达式数据类型一致
	3. case后面可以带多个表达式,使用逗号隔开.比如 case 表达式1,表达式2, ...
	4. case后面的表达式如果是常量值(字面量),则要求不能重复
	5. case后面不需要带break,程序匹配到一个case后就会执行对应的代码,然后退出switch,如果一个都匹配不到,则执行default
	6. default语句不是必须的
	7. switch后也可以不带表达式, 类似 if else 这样的分支来使用
	8. switch后也可以直接声明/定义一个变量,分号结束,不推荐
	9. switch穿透(fallthrough),如果在case语句块后增加 fallthrough,则会继续执行下一个case,也叫switch穿透
	10. Type Switch: switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的变量类型
	*/

	//  switch 用法一: 
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch key {
		case 'a','b':	// 只要满足其中一个表达式,就会执行这个case下面的语句块
			fmt.Println("输出了a或b")
		case 'c':
			fmt.Println("输出了c")
		case 'd','e','f','g':
			fmt.Println("输出了d 或e 或f 或g")
		default:
			fmt.Printf("输出了 %c \n",key)					
	} 

	//  switch 用法二:
	var age int = 21
	switch  {	// switch后也可以不带表达式, 类似 if else 这样的分支来使用
		case age == 10:
			fmt.Println("age == 10")
		case age >= 20 && age <= 30:
			fmt.Println("age >= 20 && age <= 30")	// switch也可以对范围进行判断
		default:
			fmt.Println("没有匹配到")		
	}

	//  switch 用法三(不推荐):
	switch grade := 91; {	// switch后也可以直接声明/定义一个变量,分号结束,不推荐
		case grade > 90:
			fmt.Println("成绩优秀")
		case grade >= 60 && age <= 90:
			fmt.Println("成绩及格")	 // switch也可以对范围进行判断
		default:
			fmt.Println("成绩不及格")		
	}

	// switch 使用 fallthrough 穿透, 穿透就是把下一个case代码块里面的语句也一并执行,即使不满足下一个case的条件表达式也会执行
	var num int = 10
	switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough	// 默认只能穿透一层,
		case 20:
			fmt.Println("ok2")
			fallthrough	// 默认只能穿透一层
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("没有匹配到")
	}
	/*
	上面代码运行结果为,因为用了穿透:
		ok1
		ok2
		ok3
	*/ 

	// Type Switch: switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型~ :%T \n", i)
		case int:
			fmt.Printf(" x 是 int 型 \n")
		case float64:
			fmt.Printf(" x 是 float64 型 \n")	// x 是 float64 型
		case func(int) float64:
			fmt.Printf(" x 是 func(int) 型 \n")
		case bool, string:
			fmt.Printf(" x 是 bool 或 string 型 \n")
		default:
			fmt.Printf("未知型 \n")		 
	}

}

func testFor() {

	// for循环入门,打印10次 "Hello World"
	i := 1
	for ; i <= 10; i++ {
		fmt.Println("Hello World ",i)		 
	}

	// 上面的代码可以简化为如下形式
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World ",i)		 
	}

	/*
	for循环控制注意事项和细节说明
	1.循环条件是返回一个布尔值的表达式
	2.for循环的第二种使用方式. 将 变量初始化 和 变量迭代 写到其它位置
		for  循环判断条件{
			// 循环执行语句
		}
	3.for循环的第三种使用方式.下面的写法等价 for ; ; {} ,是一个无线循环,通常需要配合break语句使用 
		for {
			// 循环执行语句
		}
	4. Golang提供 for-range 的方式,可以方便便利字符串和数组
	*/
	
	// for 循环第二种写法
	j := 1			// 循环变量初始化
	for j <= 2 {	// 循环条件
		fmt.Println("循环第二种写法-", j)
		j++			// 循环变量迭代
	}

	// for 循环第三种写法,通常会配合break使用
	k := 1			
	for  {			// 这里等价于  for ; ; {
		if k <=3 {
			fmt.Println("循环第三种写法-", k)	
		}else{
			break	// break就算跳出这个for循环
		}
		k++
	}

	// 遍历字符串,不能含中文.
	// 含有中文就会出现乱码,原因是传统方式对字符串的遍历是按照字节来遍历,而一个汉字在utf-8编码是对应三个字节.
	// 使用传统遍历方式出现中文乱码的问题解决方式是: 需要将 str 转成 []rune 切片
	str := "hello,world"
	for i:=0; i < len(str); i++{
		fmt.Printf("i=%d, val=%c \n", i, str[i])	// 使用到下标	 
	}

	// 解决乱码
	str = "hello 中国"
	str2 := []rune(str)		// 就是把 str 转成 []rune
	for i:=0; i < len(str2); i++{
		fmt.Printf("i=%d, val2=%c \n", i, str2[i])
	}

	// "for-range"方式遍历,不会出现乱码问题
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)	
	}
}

func testWhile() {
	
	// Golang没有 while 和 do...while 语法(比如 java、c等语言就有),可以通过for循环来实现其使用效果

	/*
	for循环实现 while 的效果如下:
		循环变量初始化
		for{
			if 循环条件表达式{
				break 	// 跳出for循环
			}
			循环操作(语句)
			循环变量迭代
		}

	for循环实现 do ... while 的效果如下:
		for{
			循环操作(语句)
			循环变量迭代
			if 循环条件表达式{
				break 	// 跳出for循环
			}
		}
	*/
	
	// for循环实现 while 的效果
	var i int = 1	// 循环变量初始化
	for{
		if i > 2 {	// 循环条件表达式
			break	// 跳出for循环,结束for循环
		}
		fmt.Println("用for实现while -", i)	// 循环操作(语句)
		i++			// 循环变量迭代
	} 
	fmt.Println("i = ", i)	// i =  3
	
	
	// for循环实现 do ... while 的效果
	// 下面的循环时先执行再判断,因此至少执行一次
	// 当循环条件成立后,就会执行break,break就算跳出循环,结束
	var j int = 1
	for{
		fmt.Println("用for实现 do ... while -", j)
		j++
		if j > 2{
			break
		}
	}
}