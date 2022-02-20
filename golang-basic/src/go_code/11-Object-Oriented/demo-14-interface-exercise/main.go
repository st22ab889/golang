/*
Interface 接口是Golang中内部定义的一个接口, 如下:
	type Interface interface{
		// Len is the number of elements in the collection
		Len() int
		// Less reports whether the element with
		// index i should sort before the element with index j
		Less(i, j int) bool
		// Swap swaps the elements with indexes i and j
		Swap(i, j int)
	}

要使用sort包下面的Sort函数,需要实现 Interface 这个接口,Sort函数声明如下:
	func Sort(data Interface)


接口和继承的区别:
	1) 当A结构体继承了B结构体,那么A结构体就自动继承了B结构体的字段和方法,并且可以直接使用.
	2) 当A结构体需要扩展功能,同时不希望去破坏继承关系,则可以去实现某个接口即可.因此可以认为:接口是对继承机制的补充.

接口 VS 继承:
	1) 接口和继承解决的问题不同
		继承的价值主要在于: 解决代码的复用性和可维护性
		接口的价值主要在于: 设计,设计好各种规范(方法),让其它自定义类型去实现这些方法
	2) 接口比继承更加灵活,继承是满足 is - a 的关系,而接口只需满足 like - a 的关系
	3) 接口在一定程度上实现代码解耦 
*/

package main
import (
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct{
	Name string
	Age int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero

//3.实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less方法就是决定你使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	// 按Hero的年龄从小到大排序
	return hs[i].Age < hs[j].Age
	//修改成对Name排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	// 交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//下面的一句话等价于三句话
	hs[i], hs[j] = hs[j], hs[i]
}

// 示例1: 先定义一个数组/切片,对 intSlice切片进行排序
func first () {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice) 	// 使用系统提供的排序方法
	fmt.Println(intSlice)
}


// 示例2: 对结构体切片进行排序
func second () {

	var heroes HeroSlice
	for i := 0; i < 10 ; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		//将 hero append到 heroes切片
		heroes = append(heroes, hero)
	}

	//排序前的顺序
	for _ , v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort(Sort函数定义是"func Sort(data Interface)",所以要使用Sort函数,需要实现 Interface 这个接口)
	sort.Sort(heroes)
	fmt.Println("-----------排序后------------")
	
	//排序后的顺序
	for _ , v := range heroes {
		fmt.Println(v)
	}
}

// 示例3: 变量交换的快速写法 
func three() {
	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i=", i, "j=", j) 	// i=20 j = 10
}

func main() {
	
	first ()

	second ()

	three()
}