package base

func Array4() {
	a := [2]int{}
	println(len(a), cap(a))
}

//输出结果：
//2 2

//内置函数 len 和 cap 都返回数组长度 (元素数量)。
