package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	fmt.Printf("a[0]:%p, %v\n", &a[0], a[0])
	fmt.Printf("a[1]:%p, %v\n", &a[1], a[1])
	fmt.Printf("a[2]:%p, %v\n", &a[2], a[2])

	for i, v := range a {
		fmt.Printf("i:%d v:%p, %v\n", i, &v, v)
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
			fmt.Println(v)
		}
		a[i] = v + 100
	}

	fmt.Println(a)
}
