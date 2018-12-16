// Switch statements express conditionals across many branches
package main

import "fmt"
import "time"

func main() {

	// Basic switch
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// We can use multiple expressions in one case statement
	// We havethe default case as an optional.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternative way to
	// if/else logic. Here we also show how the case expression 
	// can be non-constants.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon!")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch uses types instead of values instead of
	// values. We can use this to find out the type of an interface 
	// value.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
