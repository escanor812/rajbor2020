package rajbor2020

import "fmt"
import "math"

func Square(num int) {
	res := num * num 
	fmt.Println("Square = ", res)
}

func Cube(num int ) {

	res := num * num * num 
	fmt.Println("cube = " , res)
}

func Squareroot(num int ) {
	res := math.Sqrt(float64(num))
	fmt.Println("square root = ", res )
}