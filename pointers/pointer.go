package main

import "fmt"

func main() {
	i, j := 42, 2701

	fmt.Println("i:", &i)
	p := &i
	fmt.Println("p", p)
	fmt.Println("*p", *p)
	*p = 21
	fmt.Println("i:", i)

	p = &j
	*p = *p / 37

	fmt.Println(j)

}
