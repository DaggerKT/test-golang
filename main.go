package main

import ("fmt")

var a, b int = 1, 2

func main() {
	hellow()
}

func hellow() {
	numFolat := 3.14
	fmt.Println("int: ", a, b)
	fmt.Println("numFolat: ", numFolat)
	fmt.Println(a + b)
	fmt.Println(float64(a) * numFolat)
}