/*
字符串中常用的函数(可以查看手册或者官方编程指南):
1.按字节统计字符串长度 	len(str)
2.遍历字符串同时处理中文的问题	r:=[]rune(str)
3.字符串转整数 n,err := strconv.Atoi("12")
4.整数转字符串 str := strconv.Itoa(99)
5.字符串转[]byte	var bytes = []byte("hello Golang")
6.[]byte转字符串	str := string([]byte{97,98,99})
7.10进制转2、8、16进制		str := strconv.FormatInt(123, 2) // 2代表2进制,传入8代表8进制,传入16代表16进制
8.查找子串是否在指定的字符串中	strings.Contains("seafood", "foo")	// true
9.查找一个字符串有几个指定的字串	strings.Count("ceheese", "e")	// 4
10.不区分大小写的字符串比较(== 是区分字母大小写的)		fmt.Println(strings.EqualFold("abc", "Abc"))	// true
11.返回字串在字符串第一次出现的index值,如果没有返回-1  	index := strings.Index("NLT_abc", "abc")		// 4
12.返回字符串最后一次出现的index,如果没有返回-1		strings.LastIndex("Golang", "lang")
13.替将指定的子串替换成另外一个子串		strings.Replace("hello world bei Jing", "bei" , "Bei", n)	参数n表示替换几个,n=-1表示全部替换
14.将字符串用指定的分隔符拆分为字符串数组		strings.Split("hello,world", ",")
15.将字符串的字母进行大小写转换		strings.ToLower("GOLANG") 	strings.ToUpper("golang")
16.将字符串左右两边的空格去掉		strings.TrimSpace(" hello ")
17.将字符串左右两边指定的字符去掉	 strings.Trim("!hello ", " !")	// 将左右两边的 ! 和 " "去掉 
18.将字符串左边指定的字符去掉		strings.TrimLeft("! hello ", " !")	// 将左边 ! 和 " "去掉
19.将字符串右边指定的字符去掉		strings.TrimRight("! hello! ", " !") // 将右边 ! 和 " "去掉  
20.判断字符串是否以指定的字符串开头		strings.HasPrefix("ftp://192.168.10.1", "ftp")		// true
21.判断字符串是否以指定的字符串结束		strings.HasSuffix("NLT_abc.jpg", "abc")				// false
*/

package main
import(
	"fmt"
	"strconv"
	"strings"
)

func main() {

	stringExample_1()

	stringExample_2()
	
}


func stringExample_1(){
	// 按字节统计字符串的长度,Golang的编码统一为 utf-8(ascii的字符(字母和数字)占1个字节,汉字占用3个字节)
	str := "hello北京"
	fmt.Println("str len = ", len(str))	// 11

	// 遍历字符串,同时处理中文乱码的问题
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil{
		fmt.Println("转换错误 err = ", err)
	}else{
		fmt.Println("转换成功 n = ", n)
	}

	// 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T \n", str, str)	// str=12345, str=string

	// 字符串转[]byte
	var bytes = []byte("hello Golang")
	fmt.Printf("bytes = %v \n", bytes)	// bytes = [104 101 108 108 111 32 71 111 108 97 110 103]

	// byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str = %v \n", str)	// str = abc

	// 十进制转2、8、16进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制 = %v \n", str)	// 123对应的二进制 = 1111011
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制 = %v \n", str)	// 123对应的十六进制 = 7b

	// 查找子串是否在指定的字符串中
	b := strings.Contains("GreenFood", "Foo")
	fmt.Printf("b = %v \n", b)

	// 统计有几个子串
	num := strings.Count("ceheese", "e")
	fmt.Println("num = ", num)

	// 不区分大小写比较字符串
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b = %v \n", b)	// b = true
	// 使用"=="比较时区分大小写
	fmt.Println("结果:", "abc" == "Abc")	// 结果: false

	// 返回子串在字符串中第一次出现的index值,如果没有返回-1
	index := strings.Index("NLT_abcDdabc", "abc")
	fmt.Printf("index = %v \n", index)		// index = 4
	index = strings.Index("NLT_abcDdabc", "abcd")
	fmt.Printf("index = %v \n", index)		// index = -1
}

func stringExample_2(){

	// 返回子串在字符串中最后一次出现的index
	index := strings.LastIndex("Golang Golang", "lang")
	fmt.Printf("index = %v \n", index)	// index = 9

	// 将指定的子串替换成另外一个子串,n可以指定替换几个, 如果n=-1表示全部替换
	str1 := "hello world 北 Jing"
	str2 := strings.Replace(str1, "北" , "Bei", -1)
	fmt.Printf("str1 = %v , str2 =  %v \n", str1, str2)		// str1 = hello world 北 Jing , str2 =  hello world Bei Jing

	// 使用指定的字符作为分割标识,将一个字符串分成字符串数组
	strArr := strings.Split("hello,bei,jing", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v \n", i, strArr[i])	
	}
	fmt.Printf("strArr=%v \n", strArr)	// strArr=[hello bei jing]

	// 将字符串的字母进行大小写转换	
	str := strings.ToLower("GOLANG")
	fmt.Println("结果:", str)			// 结果: golang
	str = strings.ToUpper("golang")
	fmt.Println("结果:", str)			// 结果: GOLANG

	// 将字符串左右两边的空格去掉
	str = strings.TrimSpace(" Hello World ")
	fmt.Printf("str=%q \n", str)			// str="Hello World"

	// 将字符串左右两边指定的字符去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%q \n", str)	// str="he!llo"

	// 判断字符串是否以指定的字符串开头		
	b := strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b = %v \n", b)		// b = true

	//判断字符串是否以指定的字符串结束		
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("b = %v \n", b)		// b = false
}