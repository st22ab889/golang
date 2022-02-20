/*
时间和日期相关函数:
1.时间和日期相关函数,需要导入 time 包(time包提供了时间的显示和测量用的函数,日历的计算采用的时公历)
2.time.Time类型,用于表示时间
3.获取当前时间的方法	now := time.Now()	// now 的类型是 time.Time
4.获取到其它的日期信息
	time.Now().Year()		// 年
	time.Now().Month()		// 月  可以将返回的Month转成对应的数值	int(time.Now().Month())
	time.Now().Day()		// 日
	time.Now().Hour()		// 时
	time.Now().Minute()		// 分
	time.Now().Second()		// 秒
5.格式化日期时间
	方式1:
		now = time.Now()
		fmt.Printf("当前年月日: %02d-%02d-%02d-%02d-%02d-%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(),now.Second()) 
	方式2:
		fmt.Printf(now.Format("2021/10/29 20:00:00"))
		fmt.Println()
		fmt.Printf(now.Format("2021/10/29"))
		fmt.Println()
		fmt.Printf(now.Format("20:00:00"))
		fmt.Println()
		// "2021/10/29 20:00:00" 这个字符串的各个数字是固定的,必须这样写
		// "2021/10/29 20:00:00" 这个字符串各个数字可以自由组合,这样可以按程序需求来返回时间和日期
6.时间的常量,常量的作用可以在程序中用于获取指定时间单位的时间,比如想得到100毫秒 100 * time.Millisecond
	const (
		Nanosecond  Duration = 1						// 纳秒
		Microsecond          = 1000 * Nanosecond		// 微妙
		Millisecond          = 1000 * Microsecond		// 毫秒
		Second               = 1000 * Millisecond		// 秒
		Minute               = 60 * Second				// 分钟
		Hour                 = 60 * Minute				// 小时
	)
7.休眠	func Sleep(d Duration)
8.获取当前 unix时间戳 和 unixnano 时间戳
*/


package main
import(
	"fmt"
	"time"
	_ "strconv"
	_ "strings"
)

func main() {

	// 1. 获取当前的时间
	now := time.Now()
	fmt.Printf("now=%v, now type = %T \n", now, now)	// now=2021-10-29 19:25:07.207749 +0800 CST m=+0.003754501, now type = time.Time
	

	// 2. 通过now可以获取到年月日时分秒
	fmt.Printf("年=%v \n", now.Year())		// 年=2021
	fmt.Printf("月=%v \n", now.Month())		// 月=October 
	fmt.Printf("日=%v \n", now.Day())		// 日=29 
	fmt.Printf("时=%v \n", now.Hour())		// 时=19
	fmt.Printf("分=%v \n", now.Minute())	// 分=28
	fmt.Printf("秒=%v \n", now.Second())	// 秒=25


	// 3.1 格式化日期方式一:
	// 下面代码结果为 => 当前年月日: 2021-10-29 19:37:41
	fmt.Printf("当前年月日: %02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(),now.Second())
	// 下面代码结果为 => 当前年月日: 2021-10-29 19:38:26
	fmt.Printf("当前年月日: %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(),now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(),now.Second())
	fmt.Printf("dateStr=%v", dateStr)	// dateStr=当前年月日 2021-10-29 19:41:8


	// 3.2 格式化日期方式二:
	//返回当前年月日时分秒
	// "2021/10/29 20:00:00" 这个字符串的各个数字是固定的,必须这样写,中间的分割符可以变:
	fmt.Println(now.Format("2006-01-02 15:04:05"))	// 2021-10-29 19:57:39
	fmt.Println(now.Format("2006/01/02 15:04:05"))	// 2021/10/29 19:57:39
	fmt.Println(now.Format("2006_01_02 15:04:05"))	// 2021_10_29 19:58:06
	fmt.Println(now.Format("2006@01@02 15:04:05"))	// 2021@10@29 19:58:21
	fmt.Println(now.Format("2006Y01Y02 15:04:05"))	// 2021Y10Y29 19:58:39
	fmt.Println(now.Format("200699019902 15:04:05")) // 202199109929 19:59:05	//中间的分隔符甚至可以用数字

	// "2021/10/29 20:00:00" 这个字符串各个数字可以自由组合,这样可以按程序需求来返回时间和日期,如下
	// 返回当前年月日
	fmt.Println(now.Format("2006-01-02"))	// 2021-10-29
	// 返回当前时分秒
	fmt.Println(now.Format("15:04:05"))		// 19:50:54
	
	// Golang用2006、01、02、15、04、05这几个数字分别代表 年、月、日、时、分、秒
	fmt.Println(now.Format("2006"))	// 打印当前月
	fmt.Println(now.Format("01"))	// 打印当前月
	fmt.Println(now.Format("02"))	// 打印当前月
	fmt.Println(now.Format("15"))	// 打印当前月
	fmt.Println(now.Format("04"))	// 打印当前月
	fmt.Println(now.Format("05"))	// 打印当前月
	fmt.Println(now.Format("2006-05"))	// 打印当前时间的年和秒,可以自由组合这些代表 年、月、日、时、分、秒的数字,非常灵活
	

	// 4. 每隔 1秒/0.1秒 中打印一个数据, 打到3时就退出
	i := 0
	for{
		i++
		fmt.Println(i)
		// 休眠
		//time.Sleep(time.Second)				// 每隔1秒
		time.Sleep(time.Millisecond * 100)		// 每隔0.1秒
		if i == 3{
			break
		}
	}


	// 5. unix时间戳 和 unixnano 时间戳
	fmt.Println("Unix时间戳 = ", now.Unix())			// Unix时间戳 =  1635510040
	fmt.Println("UnixNano时间戳 = ", now.UnixNano())	// UnixNano时间戳 =  1635510040388596400
	/*
	Unix和UnixNano
		Unix将t表示为Unix时间,即从时间点January 1,1970 UTC到时间点t所经过的时间(单位秒).
		在windows下，rand.Seed(time.Now().Unix())作为种子，得出的随机数是随机的
			func (t Time) Unix() int64

		UnixNano将t表示为Unix时间,即从时间点January 1,1970 UTC到时间点t所经过的时间(单位纳秒).
		如果纳秒为单位的unix时间超出了int64能表示的范围,结果是未定义的.注意这就意味着Time零值调用UnixNano方法的话,结果是未定义的.
		在windows下，rand.Seed(time.Now().UnixNano())作为种子，得出的随机数并不随机
			func (t Time) UnixNano() int64
	*/
}