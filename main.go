package main

import "fmt"

func main() {
	fmt.Println("Hola mau")
	pin := 3.141592
	fmt.Println("Mi valor es ", pin)
	promedio()

}

func promedio() {
	var a, b float32
	fmt.Println("Dame el valor de a")
	fmt.Scan(&a)
	fmt.Println("Dame el valor de b")
	fmt.Scan(&b)
	fmt.Println(a + b)

}
