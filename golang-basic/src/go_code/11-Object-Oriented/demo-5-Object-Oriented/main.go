/*
面向对象编程应用实例:
	步骤: 1) 声明(定义)结构体,确定结构体名. 2) 编写结构体字段. 3) 编写结构体方法
*/

package main

import (
	"fmt"
)

// 示例 1
type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (student Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
	student.name, student.gender, student.age, student.id, student.score)

	return infoStr
}

func first () {
	var stu = Student{
		name : "tom",
		gender : "male",
		age : 18,
		id : 1000,
		score : 99.98,
	}
	fmt.Println(stu.say())
}

/*
示例2
1)编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2)声明一个方法获取立方体的体积,创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len float64
	width float64
	height float64
}
func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

func second () {
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volumn := box.getVolumn()
	fmt.Printf("体积为=%.2f", volumn)
}

/*
示例3: 景区门票
一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出
*/
type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice(){
	if visitor.Age >= 90 || visitor.Age <=8 {
		fmt.Println("考虑到安全，就不要玩了")
		return 
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元 \n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费 \n", visitor.Name, visitor.Age)
	}
}

func three () {
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序....")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}


func main() {

	first ()

	second ()

	three ()
}