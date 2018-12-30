package main

import "fmt"
import "os"
import "time"

func for_examples() {
	// 3 parts
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", i)
	}
	fmt.Println()
	// 3 parts
	for i, j := 0, 0; i < 10; i, j = i+1, j-1 {
		fmt.Printf("(%d,%d),", i, j)
	}
	fmt.Println()
	// 1 parts
	i := 0
	for i < 10 {
		fmt.Printf("%d,", i)
		i++
	}
	fmt.Println()
	// no parts
	i = 0
	for {
		if i < 10 {
			break
		}
		fmt.Printf("%d,", i)
		i++
	}
	fmt.Println()
	// for
	t0 := time.Now()
	for {
		if time.Since(t0) < 3*time.Second {
			continue
		}
		break
	}
	// for with labels
	some_func := func(i int) bool { return false } // i unused, but no compile error
outer:
	for idx, val := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		fmt.Printf("%v=%v\n", idx, val)
	inner:
		for i := 10; i > 0; i-- {
			if val == i {
				break outer
			}
			if some_func(i) {
				// function must be defined before usage
				break inner
			}
		}
	}
}

// function parameters may remain unused, no compile error
func foo(i, j int) {}

func main() {
	for_examples()
	os.Exit(1)
}
