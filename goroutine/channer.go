package goroutine

import (
	"fmt"
)

var intChan chan int //定义

func ChanMain() {

	intChan = make(chan int, 10) //初始化
	intChan <- 1
	fmt.Println(intChan)
	fmt.Printf("inChan 的值是%v 地址是%v \n", <-intChan, &intChan)
	//nums := <-intChan   // 注意 如此超出范围还在查 会死锁 error: all goroutines are asleep - deadlock!
	//fmt.Println(nums)
	fmt.Printf("inChan 的大小是%v 容量是%v \n", len(intChan), cap(intChan))

	strChan := make(chan string, 10)
	var str = "试"

	strChan <- str
	strChan <- "好"
	fmt.Printf("inChan 的大小是%v 容量是 %v \n", len(strChan), cap(strChan))

	mapchan := make(chan map[int]string, 5)
	mapo := make(map[int]string, 2)
	mapo[0] = "你"
	mapo[1] = "腻"
	mapchan <- mapo
	mapt := make(map[int]string, 2)
	mapt[0] = "啤"
	mapt[1] = "酒"
	mapchan <- mapt

	fmt.Printf("%v**\n***%v \n", <-mapchan, <-mapchan)

	//结构体

	alchan := make(chan interface{}, 10) //空接口 放什么都行 可以放各种 如动物  ：猫  狗
	alchan <- Dog{Age: 10, Name: "小黄", Color: "yellow"}
	alchan <- 1
	alchan <- "很好"
	//
	dog := <-alchan
	fmt.Printf("%T", dog)
	//fmt.Printf("%T", person1.Age)
	a := dog.(Dog) //给个断言 取属性
	// a :=(<-alchan).(dog)
	fmt.Printf(a.Color)
}

type Dog struct {
	Age   int
	Name  string
	Color string
}
