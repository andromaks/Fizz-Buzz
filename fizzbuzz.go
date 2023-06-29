package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        switch {
       // if the number can be divided by 15
	 case i%15 == 0:
            fmt.Println("FizzBuzz")
        // if the number can be divided by 3
	case i%3 == 0:
            fmt.Println("Fizz")
        // if the number can be divided by 5
	case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}

