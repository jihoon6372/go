// 함수
package main

import (
	"fmt"
	"strings"
)

// 함수 선언방법 1
// > 일반적인 리턴
func fun1(a int, b int) int {
	return a * b
}

// 함수 선언방법 2
// 리턴변수 미리 지정
func fun2(a int, b int) (total int) {
	total = a * b
	return
}

// 리턴값 여러개 반환
func fun3(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// arguments 여러가지
func fun4(names ...string) {
	fmt.Println(names)
}

func main() {
	num := fun1(1, 2)
	total, name := fun3("name")
	fmt.Println(num)
	fmt.Println(total, name)
	fun4("name1", "name2", "name3")
	fmt.Println(fun2(2, 2))
}
