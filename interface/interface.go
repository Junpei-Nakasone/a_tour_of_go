package main

import (
	"fmt"
	"math"
)

// interface(インタフェース)型は、メソッドのシグニチャの集まりで定義します。
// そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができます。

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Println("fの値:", f)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	fmt.Println("aの値:", a)

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
