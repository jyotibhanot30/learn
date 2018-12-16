// In Go, an array is a numbered sequence of elements of a specific len.
package main

import "fmt"

func main() {

	// That's how we declare an array.
	// The type of elements and length are part of the array's type.
	// By default an array is zero-valued, which forints means 0s.
	var a [5]int
	fmt.Println(a)

	// Setting and getting values is similar to any other language.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len keyword gets the length of an array.
	fmt.Println("len:", len(a))

	// We can use the following syntax to declare an initialize 
	// an array in one line.
	b := [5]int{1,2,3,4,5}

	fmt.Println("dcl:", b)

	// Multi-dimensional arrays can be used like this.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D:", twoD)
	for i := 0; i<2; i++ {
		for j := 0; j<3; j++ {
			fmt.Printf("%d ", twoD[i][j])
		}
		fmt.Printf("\n")

	}
}
