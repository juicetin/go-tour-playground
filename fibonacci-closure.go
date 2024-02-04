package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev := 1
    curr := 1
    counter := 1
    return func () int {
        var next int;
        switch counter {
            case 1:
                next = 1
            case 2:
                next = 1
        }

        next = prev + curr
        prev = curr
        curr = next

        counter += 1
        return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

