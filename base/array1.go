package base

import "fmt"

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func Array1() {
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}

//输出结果：
//[1 2 3 0 0] [1 2 3 4 5] [1 2 3 4 5 6] [   hello world tom]
//[1 2 0] [1 2 3 4] [0 0 100 0 200] [{user1 10} {user2 20}]

//一维数组：
//全局：
//var arr0 [5]int = [5]int{1, 2, 3}
//var arr1 = [5]int{1, 2, 3, 4, 5}
//var arr2 = [...]int{1, 2, 3, 4, 5, 6}
//var str = [5]string{3: "hello world", 4: "tom"}
//局部：
//a := [3]int{1, 2}           // 未初始化元素值为 0。
//b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
//c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
//d := [...]struct {
//name string
//age  uint8
//}{
//{"user1", 10}, // 可省略元素类型。
//{"user2", 20}, // 别忘了最后一行的逗号。
//}
