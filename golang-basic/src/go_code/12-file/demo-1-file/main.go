/*
文件
	1) 基本介绍:文件是数据源(保存数据的地方)的一种.文件最主要的作用就是保存数据.
	2) 文件在程序中是以流的形式来操作的.流:数据在数据源(文件)和程序(内存)之间经历的路径
		输入流:数据从数据源(文件)到程序(内存)的路径,就是读文件
		输出流:数据从程序(内存)到数据源(文件)的路径,就是写文件
	3) os.File(源码路径: ...\go\src\os\file.go) 封装所有文件相关操作,File是一个结构体
		func Create(name string) (file *File, err error)
		func Open(name string) (file *File, err error)
		func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
		func NewFile(fd uintptr, name string) *File
		// Pipe返回一对关联的文件对象.从r的读取将返回写入w的数据
		func Pipe() (r *File, w *File, err error)
		// 返回(提供给Open/Create等方法的)文件名称
		func (f *File) Name() string
		// Stat返回描述文件f的FileInfo类型值
		func (f *File) Stat() (fi FileInfo, err error)
		// Fd返回与文件f对应的整数类型的Unix文件描述符
		func (f *File) Fd() uintptr
		// Chdir将当前工作目录修改为f,f必须是个目录
		func (f *File) Chdir() error
		// 改变权限
		func (f *File) Chmod(mode FileMode) error
		func (f *File) Read(b []byte) (n int, err error)
		func (f *File) ReadAt(b []byte, off int64) (n int, err error)
		func (f *File) Write(b []byte) (n int, err error)
		func (f *File) WriteAt(b []byte, off int64) (n int, err error)
		func (f *File) WriteString(s string) (ret int, err error)
		// Seek设置下一次读/写的位置.offset为相对偏移量,而whence决定相对位置
		func (f *File) Seek(offset int64, whence int) (ret int, err error)
		// 把文件内容拷贝到硬盘中稳定保存
		func (f *File) Sync() (err error)
		// 关闭文件流
		func (f *File) Close() (err error)
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

// 打开和关闭文件
func first () {
	// 打开一个文件进行读操作: os.Open(name string) (*File, error)
	// 这里的file可以叫file对象,也可以叫file指针 或 file文件句柄
	file, err := os.Open("D:/GoProject/src/go_code/12-file/demo-1-file/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}	
	
	fmt.Printf("file=%v \n", file)	// file=&{0xc00007c780} ,可以看出这里的file是一个 *File 类型的指针
	
	// 关闭一个文件: File.Close()
	// 当函数退出时要及时关闭file句柄,否则会有内存泄漏
	err = file.Close()
	if err != nil {
		fmt.Println("close file err =", err)	
	}
}

// 读取文件的内容并显示在终端(带缓冲区的方式),使用 os.Open、file.Close、bufio.NewReader()、reader.ReadString函数和方法
func second () {
	file, err := os.Open("D:/GoProject/src/go_code/12-file/demo-1-file/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}
	defer file.Close()

	// 创建一个 *Reader,带缓冲,默认缓冲区为4096
	/* 源码: \go\src\bufio\bufio.go
		const (
			defaultBufSize = 4096
		) */
	reader := bufio.NewReader(file)
	for{
		str, err := reader.ReadString('\n') 	// 读到一个换行就结束
		if err == io.EOF {	// EOF 表示文件的末尾
			break
		}
		fmt.Println(str)
	}
}

// 读取文件的内容并显示在终端(使用ioutil一次将整个文件读取到内存中),这种方式适用于文件不大的情况.相关方法和函数(ioutil.ReadFile)
// func ReadFile(filename string) ([]byte error)
func three () {
	
	filePath := "D:/GoProject/src/go_code/12-file/demo-1-file/test.txt"
	content, err := ioutil.ReadFile(filePath)	// 使用 ioutil.ReadFile 一次将文件读取到位
	if err != nil {
		fmt.Println("read file err =", err)
	}
	fmt.Printf("%v", string(content))

	// 这里没有显示的Open文件,因此也不需要显示的Close文件.因为文件的Open和Close被封装到 ReadFile 函数内部
}


/*
写文件操作实例
基本介绍: func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
说明: os.OpenFile是一个更一般性的文件打开函数,它会使用指定的选项(如 O_RDONLY等)、指定的模式(如0666等)打开指定名称的文件.
	 如果操作成功,返回的文件对象可用于I/O.如果出错,错误底层类型是 *PathError
第二个参数: 文件打开模式(可以组合)
	const (
		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	)
第三个参数: 权限控制(linux)
	const (
		// 单字符是被String方法用于格式化的属性缩写。
		ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
		ModeAppend                                     // a: 只能写入，且只能写入到末尾
		ModeExclusive                                  // l: 用于执行
		ModeTemporary                                  // T: 临时文件（非备份文件）
		ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
		ModeDevice                                     // D: 设备
		ModeNamedPipe                                  // p: 命名管道（FIFO）
		ModeSocket                                     // S: Unix域socket
		ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
		ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
		ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
		ModeSticky                                     // t: 只有root/创建者能删除/移动文件
		// 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
		ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
	)

linux文件权限:
	r  ->  4
	W  ->  2
	X  ->  1

*/
// 创建新文件写入内容
func four (){

	filePath := "D:/GoProject/src/go_code/12-file/demo-1-file/test2.txt"
	// O_WRONLY 表示只有读权限, O_CREATE 表示如果不存在将创建一个新文件
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()					// 及时关闭file句柄
	str := "create new file\r\n"	 	// \r\n 表示换行
	// 这里有写文件操作,所以打开文件的模式一定要包括写模式
	writer := bufio.NewWriter(file)		//写入时，使用带缓存的 *Writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为 writer 带缓存,调用 WriteString 方法时,其实内容是先写入到缓存,需要调用Flush方法将缓存的数据真正写入到文件中,否则文件中会没有数据.
	writer.Flush()
} 

// 在已存在的文件中,用新内容覆盖旧的内容
func five () {

	filePath := "D:/GoProject/src/go_code/12-file/demo-1-file/test3.txt"
	// O_WRONLY 表示只有读权限, O_TRUNC 表示打开时清空文件
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)	
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()					// 及时关闭file句柄
	str := "new file is test3\r\n"		// \r\n 表示换行
	// 这里有写文件操作,所以打开文件的模式一定要包括写模式
	writer := bufio.NewWriter(file)		// 写入时,使用带缓存的 *Writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为 writer 带缓存,调用 WriteString 方法时,其实内容是先写入到缓存,需要调用Flush方法将缓存的数据真正写入到文件中,否则文件中会没有数据.
	writer.Flush()

}

// 在已存在的文件中,将原来的内容读出并显示在中断,并追加新内容
func six() {
	filePath := "D:/GoProject/src/go_code/12-file/demo-1-file/test3.txt"
	// O_RDWR 表示读写权限都有, O_APPEND 表示写操作时将数据附加到文件尾部 
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)	
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()		// 及时关闭file句柄

	// 这里有读文件操作,所以打开文件的模式一定要包括读模式
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {		// 如果读取到文件的末尾
			break
		}
		fmt.Print(str)
	}
	
	str := "append test\r\n"	// \r\n 表示换行
	// 这里有写文件操作,所以打开文件的模式一定要包括写模式
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)	//写入时,使用带缓存的 *Writer, WriterString 是 *Writer 的方法
	}
	// 因为 writer 带缓存,调用 WriteString 方法时,其实内容是先写入到缓存,需要调用Flush方法将缓存的数据真正写入到文件中,否则文件中会没有数据.
	writer.Flush()
}


func main() {

	first ()	// 打开和关闭文件

	second ()	// 读取文件的内容并显示在终端(带缓冲区的方式),使用 os.Open、file.Close、bufio.NewReader()、reader.ReadString函数和方法

	three ()	// 使用ioutil一次将整个文件读取到内存中,这种方式适用于文件不大的情况		

	// 使用 os.OpenFile()、bufio.NewWriter()、*Writer的WriterString 方法实现以下3个功能
	four ()		// 创建新文件写入内容

	five ()		// 在已存在的文件中,用新内容覆盖旧的内容

	six ()		// 在已存在的文件中,在原来的内容上追加内容

	
	seven ()	// 将一个文件的内容写入到另外一个文件(使用 ioutil.ReadFile / ioutil.WriteFile 完成)

	eight ()	// 判断文件是否存在

}

func seven () {
	file1Path := "D:/GoProject/src/go_code/12-file/demo-1-file/file-1.txt"
	file2Path := "D:/GoProject/src/go_code/12-file/demo-1-file/file-2.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
		return
	}
}

func eight () {
	/*
	Golang判断文件或文件夹是否存在的方法为: 使用 os.Stat() 函数返回的错误值进行判断:
	1) 如果返回的错误为 nil,说明文件或文件夹存在
	2) 如果返回的错误类型使用 os.IsNotExist() 判断为true,说明文件或文件夹不存在
	3) 如果返回的错误为其它类型,则不确定是否存在
	*/
	filePath := "D:/GoProject/src/go_code/12-file/demo-1-file/xy.txt"
	isExist, err := PathExists(filePath)
	if !isExist {
		fmt.Printf("err=%v\n", err)
	}
}

// 用一个函数实现判断文件或文件夹是否存在
func PathExists(path string) (bool, error){
	_,err := os.Stat(path)
	if err == nil {		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false, err
}
