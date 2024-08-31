package main //내가 작성할 패키지의 이름을 입력.
//컴파일을 하고 싶으면 main 패키지를 만들어야함.
import (
	"fmt"
	//"strings"
	//"lecture/module"
) //formatting 을 위한 package!

/*
var kjs string = "hi"
func multiple(a int, b int) int {  //뭔지 알겠지 ?? (a int, b int) === (a,b int)
	return a * b
}

func lenAndUpper(name string) (int,string) { //여러개 리턴값 갖기도 가능.
	return len(name) , strings.ToUpper(name)
}

func manyArgs(words ...string) int{ // 이러면 여러개의 args 를 받을 수 있어요
	return 1
}
//naked return 이 뭐지?!?!  return 에다가 변수명을 집어넣고 (변수는 선언되어 있는 상태) return 만 넣어줘도 작동함.


func lAU(name string) (leng int, result string) {
	leng = 4
	result = name
	return
}
*/

//defer 란 무엇인가 ...  func 가 끝났을떄 뭔가 시키는거. return 한 후에 실행이 된다! 흠 근데 설명하는게 다 달라서 다시 알아봐야 할듯.?

/*
	for 문에 대해서 할 얘기가 많은데 break label 언제 쓸까요?? 중첩 루프 탈출할떄 좋음.

if 조건 { 으로 쓰는데 } if 문 안에 변수를 만들 수 있음!
if val := 12+v ; val < 10 { 가능. }

switch 문. case 에다가 조건문을 달 수 있다.!! 그리고 variable expression 사용 가능함..
*/
func spadd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func main() { //package main 에 main function ㅇ ㅣ 있어야함.
	//프로그램의 시작점
	//fmt.Println("Hello world")
	//함수가 왜 대문자로 시작할까??!..
	//왜냐하면 go 에서는 함수를 export 하기 위해서는 대문자로 써야함. 이게 많이 헷갈릴수있음.

	//module.sayHello()     이러면 함수를 못 불러와~
	//module.Saybye()

	//var name string = "kjs" //와우 go 는 타입을 변수명 뒤에 적는군요
	//name1 := "kjs"
	// 변수 선언의 축약형. 타입을 알아서 찾아줌.
	//하지만 func 안에서만 가능하고 상수는 안됨.
	/*
		spadd(1,2,3,4,5,6)


		fmt.Println("loop in")
		F1:
		for a:=0;a < 30;a++ {
			fmt.Println(a)
			if a == 4 {
				break F1
			}
		}

	*/

	/* GO 에서 포인터 쓰기! C의 포인터랑 상당이 유사함을 알 수 있다... 근데 포인터 덧셈 연산 << 이건 안되네
	 
	a := 1
	b := &a // C 언어에서 int * aa = &b 이랫는데 반대로 쓰네.

	fmt.Println(&a, &b)
	*/

	// GO 의 array 에 대해서 알아봅시다! 다른 프로그래밍 언어와는 사뭇 다른 특징을 가지고 잇음 그리고 그냥 배열 출력 가능함.
	//GO 에서의 Slice 는 기본적으로 Array 임. but 길이가 없음.. 그냥 대괄호 안에를 비워주면 딥니다.
	// 크기를 지정해 놓으면 있는 범위까지 index  를 쓸 수 있지만 slice 를 쓰면 지금 있는 새로운 요소를 추가 할 때 index 를 사용할 수 없습니다.
	//slice 에 새로운 요소를 추가하기 위해서는 append() 라는 함수를 써야함. 
	
	//scs := [4]int{1,2,3,4} 이건 슬라이스가 아님.
	/*
	isthisslice := []int{}  // 뒤에 중괄호 까지 붙여야지 slice 정의가 됨.
	isthisslice = append(isthisslice,123) //append 한다고 그냥 추가되는게 아니라 slice 를 return 하기 때문에 값을 다시 넣어줘야됨.
	fmt.Println(isthisslice)
	*/

	// map 에 대해서 알아봅시다. map 은 key 와 value 로 이루어져있습니다..
	ThisisMap := map[int]string{1:"hi"}
	fmt.Println(ThisisMap)
}
