package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, upper string) {
	//함수가 끝나고 실행되는 코드
	defer fmt.Println("done")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func addTotal(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}

type person struct {
	name  string
	age   int
	hobby []int
}

func main() {
	//배열만들기
	array := []int{1, 2, 3, 4, 5}
	//struct만들기
	structure := map[string]string{"name": "jerrynim", "key": "value"}
	jerrynim := person{name: "jerrynim", age: 18, hobby: []int{1, 2}}
	for key, value := range structure {
		fmt.Println(key, value)
	}
	result := addTotal(1, 2, 3, 4, 5)
	fmt.Println(result, array, structure, jerrynim)
	fmt.Println(lenAndUpper("jerrynim"))
}
