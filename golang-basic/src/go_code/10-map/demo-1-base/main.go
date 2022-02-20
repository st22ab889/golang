/*
Map介绍: map是 key-value 数据结构,又称为字段或者关联数组.
基本语法:	var Map变量名 map[keytype]valuetype
	key可以是什么类型: 
		Golang中map的key可以是很多种类型:比如bool、数字、string、指针、channel,还可以是只包含前面几个类型的 接口、结构体、数组
		但key通常使用的类型为int,string
		注意: slic、map 还有 function 不可以, 因为这几个没法用 == 来判断
	valuetype可以是什么类型:valuetype的类型和key基本一样.通常为: 数字(整数,浮点数)、string、map、struct

map声明的举例:
	var a map[string]string
	var a map[string]int
	var a map[int]string
	var a map[string]map[string]string
	注意:声明不会分配内存,初始化需要make,分配内存后才能赋值和使用

内置函数make
	func make(Type, size IntegerType) Type
	内建函数make分配并初始化一个类型为切片、映射、或通道的对象.其第一个实参为类型,而非值.make的返回类型与其参数相同,而并非指向它的指针,其具体的结果取决于具体的类型.

map的三种用法:
	第一种:
		var cities map[string]string			// 声明,这时 map = nil
		cities = make(map[string]string, 10)	// 分配一个map
	第二种:
		var cities = make(map[string]string)	// 声明就直接make
	第三种:
		var cities map[string]string = map[string]string{"no1":"北京"}
		cities["no2"]="广州"

map的增删改查操作:
	1) map["key"] = value	// 如果key还没有就是增加,如果key存在就是修改
	2) delete(map, "key")	// delete是一个内置函数,如果key存在,就删除该key-value,如果key不存在就不操作,也不会报错
		map删除细节说明:
			如果要删除map的所有key,没有一个专门的方法一次删除,可以遍历key,逐个删除.
			或者 map = make(...),make一个新的,让原来的称为垃圾,被GC回收
	3) map查找

map遍历:map的遍历使用 for-range 的结构遍历	

map的长度: 使用内置函数 len
*/

package main

import (
	"fmt"
)


func first () {
	// map的声明
	var a map[string]string
	// 在使用map前需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "java"
	a["no2"] = "C/C++"
	a["no3"] = "Golang"
	a["no4"] = "kubernetes"
	fmt.Println("a = ", a)
	
	/*
	注意事项:
	1) map在使用前一定要 make,make是内置函数
	2) map的key不能重复,如果重复了,以最后一个key-value为准. map的value是可以相同的(也就说说value可以重复)
	3) map的 key-value 是无序
	*/

}

// map的三种用法
func second () {
	//第一种方式: 声明map后需要make,make的作用就是给map分配数据空间
	var cities map[string]string
	cities = make(map[string]string, 10)
	cities["c1"] = "北京"
	cities["c2"] = "广州"
	fmt.Println("cities = ", cities)

	//第二种方式
	//var dev = make(map[string]string)
	dev := make(map[string]string)
	dev["d1"] = "java"
	dev["d2"] = "kubernetes"
	fmt.Println("dev = ", dev)

	//第三种方式
	//var key map[string]string = map[string]string{
	key := map[string]string{
		"k1" : "qwertyuiop",
		"k2" : "asdfghjkl",
	}
	key["k3"] = "zxcvbnm"	// map的增加操作
	key["k1"] = "aaabbbccc"	// map的修改操作
	fmt.Println("key = ", key)

	// map查找,如果key这个map中存在"k2"这个key,那么res就会返回true,否则返回false
	val, res := key["k2"]
	if res{
		fmt.Printf("有 k2 这个 key, 值为 %v \n", val)
	} else {
		fmt.Printf("没有 k2 这个 key \n")
	}

	delete(key, "k3")		// map的删除操作
	delete(key, "k5")		// delete指定的key不存在时就不操作,也不会报错
	fmt.Println("key = ", key)

	// 如果希望一次性删除所有的key
	// 1. 遍历所有的key逐一删除;	2.直接make一个新的空间
	key = make(map[string]string)
	fmt.Println("key = ", key)

}

func three () {

	// 1. 使用 for-range 遍历map
	cities := make(map[string]string, 10)
	cities["c1"] = "北京"
	cities["c2"] = "广州"
	for k,v := range cities {
		fmt.Printf("k=%v , v=%v \n", k, v)
	}
	fmt.Println("cities 有 ", len(cities), "对 key-value")


	// 2. 使用 for-range 遍历嵌套的map
	peopleMap := make(map[string]map[string]string)
	peopleMap["p1"] = make(map[string]string, 3)
	peopleMap["p1"]["name"] = "Tom"
	peopleMap["p1"]["age"] = "20"
	peopleMap["p1"]["addr"] = "北京"

	peopleMap["p2"] = make(map[string]string, 3)
	peopleMap["p2"]["name"] = "Lily"
	peopleMap["p2"]["age"] = "18"
	peopleMap["p2"]["addr"] = "广州"

	for k1, v1 := range peopleMap {
		fmt.Println("k1 = ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v , v2=%v \n", k2, v2)
		}
	}

}

func main() {

	first ()

	second ()

	three ()
}