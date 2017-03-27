package testdefer

import "fmt"

func init() {
	/*defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")*/
	forDefer()
	forDefer1()
}

func forDefer() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()// () 表示执行 已经讲过了
	}
}

func forDefer1() {
	for i := 0; i < 5; i++ {
		defer func(i int) {//参数传递是值的考呗
			fmt.Println(i)
		}(i)// () 表示执行 已经讲过了
	}
}