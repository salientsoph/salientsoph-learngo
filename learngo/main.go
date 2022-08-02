package main

import (
	"fmt"
	"time"
)

//goRoutine: 프로그램이 작동하는 동안만 유효 (main 함수가 실행되는 동안만)
//메인함수는 goRoutine을 기다려주지않음..
func main() {

	//make(채널이름 채널로보낼종류)
	//어떤 타입의 정보를 보낼건지
	//channel을 통해 메인함수와 정보를 주고받음
	//
	c := make(chan bool)

	people := [2]string{"nico", "flynn"}
	for _, person := range people {

		//isSexy는 c라는 인수를 받음
		go isSexy(person, c) //이 함수가 끝나면 true 값을 channel을 통해 보내고 싶음
	}
	/*
		go sexyCount("nico")
		go sexyCount("flynn")
	*/
	//result := <-c //데이터 수신
	//fmt.Println(result)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}

//goRoutine: 다른 함수와 동시에 실행시키는 함수
/*
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second * 11)
	}
}
*/

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true //데이터 송신
}

//channel: goRoutine과 메인함수 사이에 정보를 전달하기 위한 방법
