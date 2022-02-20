/*
 多维数组 - 二维数组
	使用方式1:先申明/定义,再赋值
		语法:  	var 数组名 [大小][大小]类型
				var arr [2][3]int[][]
	使用方式2:直接初始化
		声明:	var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
		赋值(有默认值,比如int类型的默认值就是0)
	
	二维数组在声明/定义时也对应有四种写法(和一维数组类似)
		var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
		var 数组名 [大小][大小]类型 = [...][大小]类型{{初值...},{初值...}}
		var 数组名 = [大小][大小]类型{{初值...},{初值...}}
		var 数组名 = [...][大小]类型{{初值...},{初值...}}

	二维数组的遍历:
		1) 双层for循环遍历
		2) for-range方式遍历
*/

package main
import(
	"fmt"
)

func first() {
	
	/*
	二维数组案例演示,打印如下图形:
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/

	// 定义/声明二维数组
	var arr [4][6]int	// [4][6]表示4行六列
	// 赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3	
	fmt.Println(arr)
	// 遍历二维数组,按照要求输出图形
	for i:=0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

}

func second() {
	
	// 二维数组定义时直接初始化
	var arr = [4][6]int{{0, 0, 0, 0, 0, 0},{0, 0, 1, 0, 0, 0},{0, 2, 0, 3, 0, 0},{0, 0, 0, 0, 0, 0}}
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3	

	// 二维数组的实质是: 一维数组中的每个元素又是数组
	fmt.Printf("arr[0]的地址 = %p \n", &arr[0])			// arr[0]的地址 = 0xc00010a0c0
	fmt.Printf("arr[0][0]的地址 = %p \n", &arr[0][0])	// arr[0][0]的地址 = 0xc00010a0c0

	fmt.Printf("arr[1]的地址 = %p \n", &arr[1])			// arr[1]的地址 = 0xc00010a0f0
	fmt.Printf("arr[1][0]的地址 = %p \n", &arr[1][0])	// arr[1][0]的地址 = 0xc00010a0f0	
}

func three() {
	var arr = [2][3]int{{1, 2, 3},{4, 5, 6}}

	// for循环遍历
	for i:=0; i < len(arr); i++ {
		for j:=0; j < len(arr[i]); j++ {
			fmt.Printf("%v \t", arr[i][j])
		}
	}

	fmt.Println()

	// forr-range遍历二维数组
	for i, v := range arr {
		for j , v2 := range v {
			fmt.Printf("arr[%v][%v]=%v \t", i, j, v2)
		}
	}
}


func main() {

	first() 

	second()

	three()
}



