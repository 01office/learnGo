package main

import "fmt"

func main() {
	fmt.Println("Hi, everyone!")

	var i, j, tmp int
	var a [10]int
	fmt.Println("Please enter 10 numbers: ")
	for i = 0; i < 10; i++ {
		fmt.Scanln(&a[i])
	}

	for i = 0; i < 10; i++ {
		for j = 1; j < 10 - i; j++ {
			if a[j - 1] > a[j] {
				tmp = a[j]
				a[j] = a[j - 1]
				a[j - 1] = tmp
			}
		}
	}

	fmt.Println("After sorted:")
	for i = 0; i < 10; i++ {
		fmt.Println(a[i], " ")
	}
}

