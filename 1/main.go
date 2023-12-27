package main

//Создай внешний пакет с именем mymath, который будет содержать функциональность для выполнения математических операций с версией v1.0.0

import (
	"fmt"
	"github.com/MikeBazh/mymath"
)

func main() {
	x := 4.0
	y := 2.0
	fmt.Println(mymath.Abs(x))
	fmt.Println(mymath.Acos(x))
	fmt.Println(mymath.Acosh(y))
}
