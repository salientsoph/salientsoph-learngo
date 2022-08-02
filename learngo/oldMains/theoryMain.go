package oldMains

import (
	"fmt" //package for formatting
	"strings"

	"github.com/salientsoph/learngo/something"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func canIDrink2(age int) bool {
	/*switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false*/

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func canIDrink(age int) bool {
	//koreanAge := age + 2 -> 이것과 동일

	//if-else에서만 쓰기 위한 변수 정의
	if koreanAge := age + 2; koreanAge < 18 { //variable expression (if를 선언하면서 변수 선언 가능)
		return false
	} //if 문에서 return으로 끝나면 else는 필요 x
	return true
}

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

func theoryMain() {
	fmt.Println("Hello world!")
	something.SayHello() //exported function (대문자 사용)
	//something.sayBye() //사용불가. private function

	fmt.Println("----------------------------1.3 상수, 변수")

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

	fmt.Println("----------------------------1.4 functions(1)")

	fmt.Println("multiply: ", multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	//return 값이 2개면 무조건 2개를 받아야함
	fmt.Println("totalLength, upperName: ", totalLength, upperName)

	//무시하고 싶다면
	totalLength2, _ := lenAndUpper("new value")
	fmt.Println("totalLength2, upperName: ", totalLength2)

	fmt.Print("파라미터 여러개 받기: ")
	repeatMe("nico", "lynn", "dal", "straw", "tiger")

	fmt.Println("----------------------------1.5 functions(2)")

	totalLength3, upperName3 := lenAndUpper2("nico")
	fmt.Println("totalLength3, upperName3: ", totalLength3, upperName3)

	fmt.Println("----------------------------1.6 if with a twist")

	fmt.Println(canIDrink(15))

	fmt.Println("----------------------------1.7 switch")

	fmt.Println(canIDrink2(52))

	fmt.Println("----------------------------1.8 pointers!")

	a := 2
	b := &a
	a = 10
	*b = 2022 //b를 통해서 a값 변경
	fmt.Println(a, *b)
	fmt.Println(&a, &b) //메모리 주소

	fmt.Println("----------------------------1.9 arrays and slices")

	//arrays
	names := [5]string{"nico", "lynn", "dal"} //5개의 요소를 갖는 array
	names[3] = "alla"
	names[4] = "alla"
	//names[5] = "alla" //array index 5 out of bounds [0:5]
	fmt.Println(names)

	//slices
	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "lala") //append(slice, new value): 새로운 값이 추가된 slice를 return함
	//변수에 담아준다
	fmt.Println(names2)

	fmt.Println("----------------------------1.10 maps")

	//map[key]value
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}

	for _, value := range nico {
		fmt.Println(value)
	}

	fmt.Println("----------------------------1.11 structs")
	//structs == object
	//map 보다 flexible

	/*
		favFood := []string{"kimchi", "ramen"}
		hill := person{"hill", 18, favFood}
		fmt.Println(hill.favFood)*/

	//key, value 를 모두 적어주거나 value만 적는다. 몇개는 key, value 적고 몇개는 value 적는 것 불가
	favFood := []string{"kimchi", "ramen"}
	hill := person{name: "hill", age: 18, favFood: favFood}
	fmt.Println(hill.favFood)

}
