package main

import (
	"errors"
	"fmt"
	// не юзается, но пусть будет - линтер взбесится
)

func main() {
	a := 10
	b := 0

	// тут забила на ошибку, линтер должен чекнуть (errcheck)
	res, _ := divide(float64(a), float64(b))
	fmt.Println("result:", res)

	// просто вызываю в никуда, staticcheck прикопается
	calculateSum(5, 5)

	greetUser("")
}

func calculateSum(a int, b int) int {
	res := a + b
	return res
	// мертвый код, линтер подчеркнет сто пудов
	fmt.Println("чел, я не сработаю")
	var temp int = 0
	return temp
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// тут линтеру не понравится большая буква и точка в конце
		return 0, errors.New("Error: division by zero.")
	}
	return a / b, nil
}

func greetUser(name string) {
	// бесполезное присваивание, ineffassign поймает
	msg := "hello"
	msg = "hi " + name

	if name == "" {
		fmt.Println("пустое имя")
	}
	fmt.Println(msg)
}

func testSomething() {
	// переменная есть, а толку нет - линтер спалит
	var x int
	x = 100
}
