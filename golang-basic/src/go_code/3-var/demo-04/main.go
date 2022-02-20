
/*
Golang中没有专门的字符类型,如果要存储单个字符(字母),一般使用byte来保存

字符串就是一串固定长度的字符连接起来的字符序列
Go的字符串是由单个字节连接起来的
也就是说对于传统的字符串是由字符组成的,而Go的字符串不同,它是由字节组成的.
官方将string归属到基本数据类型:	https://tour.go-zh.org/basics/11
*/

package main
import(
	"fmt"
)

func main()  {
	
	// 1. 如果保存的字符在 ASCII表的,比如[0-1,a-z,A-Z..],直接保存到byte
	// 2. 如果保存的字符对应码大于255,可以考虑使用int保存
	// 3. 如果需要字符的方式输出,需要用到格式化输出, 比如 fmt.Printf("c1 = %c", c1) 

	var c1 byte = 'a'	// 字符只能使用单引号
	var c2 byte = '0'	// 字符的0

	// 当直接输出byte值,就是输出了字符对应的码值
	fmt.Println("c1 = ", c1)	// c1 =  97
	fmt.Println("c2 = ", c2)	// c2 =  48

	// 如果希望输出对应字符,需要使用格式化输出
	fmt.Printf("c1 = %c, c2=%c \n", c1, c2)	// c1 = a, c2=0 

	// 编译时报"constant 22909 overflows byte", 说明'好'的码值是22909,这是一个Unicode的十进制码，byte能表示的范围是"0~255"
	// var c3 byte = '好'	

	// int 能表示范围可以比22909大,所以用int没问题
	var c3 int = '好'
	fmt.Printf("c3 = %c, c3对应的码值=%d", c3, c3)	// c3 = 好, c3对应的码值=22909

	// '好'对应的16进制Unicode编码是: \u597D 
	//  \u597D 意思是: \表示转义, \u 表示是unicode码, 一般其后跟4个16进制数,因此,一般为unicode-16
	// 	类似: 
	//		\x：只是 16 进制的意思，后边跟两位，则表示单字节编码；
	//		\d：十进制；
	//		\o：八进制；
	//	转义及编码（\u, \x）: https://www.cnblogs.com/videring/articles/7182020.html
	
	/*
	 	16进制Unicode编码（可以同时用于HTML与JS中）,如下:
			<!--应用在HTML中示例-->
			<div>\u597D</div>
			<script>
			// 应用在JavaScript中示例
			alert('\u597D');
			</script>

		'好'对应的16进制Unicode编码是 597D, 十进制Unicode编码是 22909, 在网页中还有如下用法:
				<div>&#x597D;</div>
				<div>&#22909;</div>
			&#x 或 &# 是 HTML、XML 等 SGML 类语言的转义序列(escape sequence), &#开头的接十进制数字, &#x开头的接十六进制数字

		Unicode 编码&解码: https://www.css-js.com/tools/unicode.html
	*/

	/*
	Unicode 和 UTF-8 之间的关系和区别(https://www.zhihu.com/question/23374078):
		UTF-8 是一种 Unicode 编码形式(Encoding Form),建立了一个 Unicode 标量值(Scalar Value)和字节序列(Byte Sequence)的一个真子集之间的双射(Bijection)
		utf8是对unicode字符集进行编码的一种编码方式
	*/

	charExample();

}

func charExample(){

	/*
	字符类型使用细节
	1. 字符常量是用单引号('')括起来的单个字符.例如: var c1 byte = 'a'
	2. GO中允许使用转义字符'\'来将其后的字符转变为特殊字符型常量。 例如: var c3 char = "\n"
		\a		响铃
		\b		退格
		\f		换页
		\n		换行
		\r		回车
		\t		制表符
		\v		垂直制表符
		\'		单引号(只用在 '\" 形式的 rune 符号面值中)
		\"		双引号(只用在"..."形式的字符串面值中)
		\\		反斜杠
	3. Go语言的字符使用UTF-8编码。英文字母占1个字节,汉字占3个字节 
		查询字符对应的utf8码值: http://www.mytju.com/classcode/tools/encode_utf8.asp
	4. 在Go中,字符的本质是一个整数,直接输出时,是该字符对应的UTF-8编码的码值
	5. 可以直接给某个变量赋一个数字,然后格式化输出时 %c,会输出该数字对应的unicode字符
	6. 字符类型是可以进行运算的,相当于一个整数,因为它都对应有Unicode码
	*/

	var c4 int = 22269	// "国"对应的Unicode码是22269
	fmt.Printf("c4 = %c \n", c4)	// c4 = 国
	
	var n1 = 10 + 'a'	// "a"的码值是97. (如果字符的码值属于[0,255]这个区间,那么这个字符的Unicode、UTF-8、ASCII码都是一样的)
	fmt.Println("n1 = ", n1)	 // n1 =  107
}