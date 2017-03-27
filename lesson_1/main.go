package main

import (
	//_ "csdn_go/lesson_1/closure"
	_ "csdn_go/lesson_1/testdefer"
)

func main() {
	/*fmt.Println(add)
	fmt.Println(add(5, 5))
	fmt.Println(change)
	fmt.Println(change(4, 5))
	a := add
	fmt.Println(a(5, 6))

	var b func(int, int) int = func(a, b int) int {
		return a + b
	}
	fmt.Println(b(5, 6))

	var c = func(a, b int) int {
		return a + b
	}
	fmt.Println(c(5, 6))

	d := func(a, b int) int {
		return a + b
	}
	fmt.Println(d(5, 6))

	fmt.Println(func(a, b int) int {
		return a + b
	}(5, 6))*/
}
/**
相加
 */
func add(x int, y int) int {
	//这里的a可以省略
	return x + y
}
/**
交换
 */
func change(x int, y int) (int, int) {
	return y, x
}

/**
可以申明多个同一类型的值
 */
func add1(x, y int) int {
	return x + y
}
/**
可以直接把返回值类型变量当成局部变量，并且已经是声明好的，那么我们可以直接return，因为go已经知道你要返回谁
 */
func change1(x, y int) (a, b int) {
	a, b = y, x;
	return
}