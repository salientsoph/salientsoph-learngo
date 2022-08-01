package main

import (
	"fmt" //package for formatting

	"github.com/salientsoph/learngo/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello() //exported function (대문자 사용)
	//something.sayBye() //사용불가. private function
}
