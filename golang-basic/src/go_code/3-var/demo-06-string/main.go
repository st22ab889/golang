
package main
import(
	"fmt"
)

func main()  {
	
	// 字符串就是一串固定长度的字符连接起来的字符序列
	// Go的字符串是由单个字节连接起来的。go语言的字符串的字节使用UTF-8编码标识Unicode文本

	var country string = "中华人民共和国"
	fmt.Println(country)

	// 使用细节
	// 1.Go语言的字符串的字节使用UTF-8编码标识Unicode文件,这样Golang统一使用UTF-8编码,不会有乱码问题
	// 2.Go语言中字符串不可变,一旦赋值不能修改
	// 		country[0]='a'	// 不能修改
	// 3.字符串两种表现形式
	//		(1) 双引号,会识别转义字符
	//		(2) 反引号,以字符串的原生形式输出,包括换行和特殊字符,可以实现防止攻击、输出源代码效果
	// 4.字符串拼接方式
	// 5.当一行字符串太长时,需要使用到多行字符串

	// 使用双引号
	str := "中华人民\n共和国"		
	fmt.Println(str)
	
	// 使用反引号 `` ,输出源代码效果
	str2 :=`
	package main
	import(
		"fmt"
	)
	`
	fmt.Println(str2)	

	// 使用 + 或 += 实现字符串拼接
	var str3 string
	str3 = "hello" + " world"
	fmt.Println(str3)
	str3 += " !"
	fmt.Println(str3)

	// 当一行字符串太长时,需要使用到多行字符串
	str4 := "hello" +		// + 号放在此行,放在下一行会报错
		" world !"
	fmt.Println(str4)

}


