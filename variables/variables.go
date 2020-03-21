// 변수와 상수
package main

import "fmt"

func main() {
	// 상수
	const name string = "name"
	fmt.Println(name)

	// 변수
	var name2 string = "name 2"
	fmt.Println(name2)

	// 타입은 축약형으로 사용 가능
	// 축약형으로 사용 할 경우 초기에 할당되는 값에 따라서 타입을 자동으로 결정해줌
	name3 := "name 3" // string
	num := 1          // int

	fmt.Println(name3)
	fmt.Println(num)
}
