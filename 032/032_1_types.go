package main

import "fmt"

var a int // 미리지정
var d = "Nevermind"
var e string = "Nevermind" // 동적 프로그래밍언어가 아님.

func main() {
	fmt.Println("a -> ", a)
	fmt.Printf("a Type -> %T\n", a)

	b := 3 // 타입 지정을 안한다면? 1
	fmt.Println("b -> ", b)
	fmt.Printf("b Type -> %T\n", b)

	c := 'a'                        // 타입 지정을 안한다면? 2 어라?
	fmt.Println("c -> ", c)         // c ->  97
	fmt.Printf("c Type -> %T\n", c) // c Type -> int32

	fmt.Println("d -> ", d)
	fmt.Printf("d Type -> %T\n", d)

	fmt.Println("e -> ", e)
	fmt.Printf("e Type -> %T\n", e)

}
