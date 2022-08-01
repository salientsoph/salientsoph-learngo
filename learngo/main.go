package main

import (
	"fmt" //package for formatting
	"strings"

	"github.com/salientsoph/learngo/something"
)

func lenAndUpper2(name string) (length int, upperCase string) {
	defer fmt.Println("I'm done") //defer: function이 끝났을 때 작동
	length = len(name)
	upperCase = strings.ToUpper(name)
	return //naked return (명시하지 않아도됨)
}

//파라미터 무제한으로 받기 -> type 앞에 ... 사용
func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a int, b int) int { //func multiply(a , b int) int { //-> a, b 모두 int로 인식
	return a * b
}

func main() {
	fmt.Println("Hello world!")
	something.SayHello() //exported function (대문자 사용)
	//something.sayBye() //사용불가. private function

	fmt.Println("----------------------------상수, 변수")

	//상수: 변수지만 값을 바꿀 수 없음 -> const
	//변수: -> let
	const name string = "nico"
	//name = "Lynn"  //const라 바꿀 수 없음
	fmt.Println("name: ", name)

	var name2 string = "nico2"
	name2 = "lynn"
	fmt.Println("name2: ", name2)

	name3 := "nico3" //== var //func 안에서는 이렇게도 가능
	//type은 알아서 정해줌 (type은 바꿀 수 없음)
	fmt.Println("name3: ", name3)

	fmt.Println("----------------------------functions(1)")

	fmt.Println("multiply: ", multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	//return 값이 2개면 무조건 2개를 받아야함
	fmt.Println("totalLength, upperName: ", totalLength, upperName)

	//무시하고 싶다면
	totalLength2, _ := lenAndUpper("new value")
	fmt.Println("totalLength2, upperName: ", totalLength2)

	fmt.Print("파라미터 여러개 받기: ")
	repeatMe("nico", "lynn", "dal", "straw", "tiger")

	fmt.Println("----------------------------functions(2)")

	totalLength3, upperName3 := lenAndUpper2("nico")
	fmt.Println("totalLength3, upperName3: ", totalLength3, upperName3)

}
