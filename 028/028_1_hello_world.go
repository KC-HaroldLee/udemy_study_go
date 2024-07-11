package main

import (
	"fmt"
)

func main() { // 빨간 줄이 뜨기는 한다.
	fmt.Println("this is main()")
	foo()
}

func foo() {
	fmt.Println("this is foo()")
}
