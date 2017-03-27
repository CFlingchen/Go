package closure

import "fmt"

func init() {
	f := closure(10);
	fmt.Println(f(11))
	fmt.Println(f(12))
}

/**
测试闭包
 */
func closure(a int) func(int) int {
	fmt.Printf("%p\n", &a)
	return func(b int) int {
		fmt.Printf("%p\n", &a)
		return a + b;
	}
}
