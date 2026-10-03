package main

import (
	"fmt"
	"time"
)

// main: start doing work
func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 22:
		fmt.Println("It's after 22h")
	default:
		fmt.Println("It's after noon, but before 22h")
	}

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool:", t)
		case int:
			fmt.Println("I'm an int:", t)
		case string:
			fmt.Println("I'm a string:", t)
		case float64:
			fmt.Println("I'm a float64:", t)
		default:
			fmt.Printf("Don't know type %T: %v", t, t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(1.0)
	whatAmI("hey")
	whatAmI([]string{"hello", "world"})
}
