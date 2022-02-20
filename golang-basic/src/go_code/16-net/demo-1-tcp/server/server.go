package main

import (
	"fmt"
	"net"	// 做网络socket开发时,net包含有所需要的所有方法和函数
	_ "io"
)


func process(conn net.Conn) {
	
	defer conn.Close() 		 //关闭conn

	//这里循环接收客户端发送的数据
	for {
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())

		buf := make([]byte, 1024)		//创建一个新的切片
		// 等待客户端通过conn发送信息,如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		n, err := conn.Read(buf)  // 收到客户端发送的信息就从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return 
		}

		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n])) 
	}
}


func main() {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")	// tcp 表示使用网络协议是tcp, 0.0.0.0:8888 表示在本地监听 8888端口
	if err != nil {
		fmt.Println("listen err=", err)
		return 		
	}
	defer listen.Close() 	// 延时关闭listen

	//循环等待客户端来链接
	for {
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		}else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		// 准备其一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen suc=%v\n", listen)
}