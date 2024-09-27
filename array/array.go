package main

var floatArr = [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}

func main() {
	for i := 0; i < len(floatArr); i++ {
		println(floatArr[i])
	}
}