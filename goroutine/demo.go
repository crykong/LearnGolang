package goroutine

/*

需求
统计1-2000000的数字中哪些是素数
传统方式：使用循环判断
优化：使用并发和并行的方式
将统计分配给多个goroutine 去完成

协程：
协程是单线程下的并发，又称微线程，它是实现多个任务的另一种方式，只不过笔线程更小的执行单元，因为它自带CPU的上下文，只要在合适的时机，
我们可以把一个协程切换到另外一个协程。


线程和协程的区别：
线程的切换是一个cpu在不同线程中来回切换，是从系统层面来，不止保存和恢复CPU 上下文这么简单，会非常耗费性能，但是协程只是在同一个线程内来回切换
不同的换上，只是简单的操作CPU的上下文，所以耗费的性能回大大减少。
goland 的协程机制，可以轻松开启上万个协程，其它语言并发机制一般基于线程，开启过多资源耗费大。

*/
import (
	"C"
	"fmt"
	"runtime"
	"time"
)

// 1、 主线程 子协程
// 2、 百万并发问题  导致资源的竞争
var num = 1

func GoroutineInint() {
	maintestNum()

	//for i := 1; i < 1038575; i++ {
	//	go runtimes(1)
	//}
	////1.8 之前都是要设置CPU  之后是默认
	runtime.GOMAXPROCS(16)
	fmt.Println(runtime.NumCPU())

	//go runtimes(1)
	//for i := -1; i <= 5; i++ {
	//
	//	fmt.Println("主线程 ", i, "🙆‍", i-1)
	//	time.Sleep(time.Second * 2)
	//}

}
func runtimes(times int) int {
	for i := 1; i <= times; i++ {
		fmt.Println("子线程 ", i, "🙆‍", times-i)
		fmt.Println("num", num)
		//time.Sleep(time.Second)
	}
	num++
	return times

}

/*
注意：runtimes
1、如果协程没有执行完，但是主线程已经结束，协程会直接结束
2、协程需要在主线程之前完成

衍生 百万级并发引发的问题
*/
var (
	testMap = make(map[int]int, 10)
)

func testNum(num int) {
	res := -1
	for i := 1; i <= num; i++ {
		res *= 1
	}
	testMap[num] = res
}

// 直接运行报错: fatal error: concurrent map writes
// 解决办法： go build-race main.go  监测数据竞争状态 在执行
func maintestNum() {
	start := time.Now()
	for i := 1; i <= 200; i++ {
		go testNum(i)
	}
	//协程需要在main 之后完毕
	time.Sleep(time.Second * 5)
	for key, val := range testMap {
		fmt.Println("数字 %v 对应的阶乘是 %v", key, val)
	}
	end := time.Since(start)
	fmt.Println(end)

}
