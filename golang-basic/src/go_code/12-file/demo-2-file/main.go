/*
拷贝文件: 将一个文件拷贝到另外一个目录下
io包: func Copy(dst Writer, src Reader) (written int64, err error)
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

// 拷贝文件
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return 
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func first () {
	srcFile := "D:/GoProject/src/go_code/12-file/demo-2-file/image.png"
	dstFile := "D:/GoProject/src/go_code/12-file/demo-2-file/temp/image2.png"
	//调用CopyFile 完成文件拷贝
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}


func main() {

	first ()	// 拷贝文件

	second ()	// 统计一个文件中含有的英文、数字、空格及其它字符数量
}

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount int 	// 记录英文个数
	NumCount int 	// 记录数字的个数
	SpaceCount int 	// 记录空格的个数
	OtherCount int 	// 记录其它字符的个数
}

// 统计英文、数字、空格和其它字符的数量
func second () {
	
	// 打开一个文件, 创建一个Reader, 每读取一行就统计该行有多少个英文、数字、空格和其他字符, 然后将结果保存到一个结构体
	fileName := "D:/GoProject/src/go_code/12-file/demo-2-file/main.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	
	//定义个CharCount 实例
	var count CharCount

	//创建一个Reader,然后开始循环读取fileName的内容
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { 		//读到文件末尾就退出
			break
		}
		//遍历 str ，进行统计
		for _, v := range str {
			
			switch {
				case v >= 'a' && v <= 'z':
						fallthrough 		//穿透
				case v >= 'A' && v <= 'Z':
						count.ChCount++
				case v == ' ' || v == '\t':
						count.SpaceCount++
				case v >= '0' && v <= '9':
						count.NumCount++
				default :
						count.OtherCount++
			}
		}
	}

	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}