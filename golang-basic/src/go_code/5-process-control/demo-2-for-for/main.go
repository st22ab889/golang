
/*
多重循环控制
1.将一个循环放在另一个循环体内,就形成了嵌套循环.在外边的for称为外层循环,在里面的for循环称为内存循环.(建议一般使用两层,最多不要超过三层)
2.实质上,嵌套循环就是把内层循环当成外层循环的循环体.当只有内层循环的循环条件为false时,才会完全跳出内层循环,才可结束外层的当次循环,开始下一次循环
3.外层循环次数为m次,内层为n次,则内层循环实际上需要执行 m * n 次
*/

package main
import(
	_ "fmt"
	_ "unsafe"
	_ "strconv"
)

func main() {

}