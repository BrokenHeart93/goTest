package base

import (
	"fmt"
)

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func Array3() {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)

	test(a)
	fmt.Println(a)
}

//输出结果：
//a: 0xc000018080
//x: 0xc000018090
//[0 0]

//值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
