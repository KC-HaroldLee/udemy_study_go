package main

import "fmt"

// x := 42 //밖에 있으면 에러 < syntax error: non-declaration statement outside function body
var y = 43 //밖에 있어도 됨

// 타입만 지정하는 것도 된다.
var z int

func main() {
	// 단축 선언 연산자
	x := 42 // 블록 안에 있어야함
	fmt.Println("x:", x)

	// var 사용
	// var y = 43  //밖에 있어도 됨
	// "범위를 제한하세요!""
	// 따라서 특별한 경우 아니면 var를 쓰지 않는다?
	fmt.Println("y:", y)

	// 이건 되겠지만
	z = 1
	// z = 0.1 // << 에러

	// var t int  // 선언에러
	fmt.Println("z:", z)
	foo()
}

func foo() {
	fmt.Println("foo in ", y)
}
