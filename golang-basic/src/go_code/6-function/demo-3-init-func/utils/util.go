
package utils
import(
	"fmt"
)

// Age 和 Name 是全局变量,需要在其它源文件中使用,但是需要先初始化Age和Name
var Age int
var Name string



func init ()  {
	Age = 100
	Name = "test"
	fmt.Println("Util -> init()")
}