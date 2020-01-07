package main

import . "fmt"
import "strconv"
import "errors"

var err_fmt = errors.New("format error")
var err_parse error = errors.New("parse error")
var err_unimpl = errors.New("unimplemented error")

type Point struct {
	x, y int
}

func ParsePoint(v string) (Point, error) {
	// return Point{x: 0, y: 0}, err_unimpl
	return Point{x: 0, y: 0}, nil
}

func test5() {
	point, err := ParsePoint("")
	if err == nil {
		// Printf("no error, p = (%v,%v)\n", point.x, point.y)
		Printf("no error, p = (%d,%d)\n", point.x, point.y)
	} else {
		Println(err)
	}
}

var x int
var y int = 1

// z:= 1
func test4() {
	z := 1
	Z := 1
	println(z, Z)
}

func test1() {
	x, _ := strconv.ParseBool("true")
	Println(x) // true
	x, _ = strconv.ParseBool("false")
	Println(x) // false
	x, _ = strconv.ParseBool("trues")
	Println(x) // false
	x, err := strconv.ParseBool("trues")
	Println(x, err) // false
}

func test2() {
	a := 1
	// a := 2 // error, no new variables on the left side
	a, b := 2, 3 // no error, even a is already defined
	Println(a, b)

	x := 1
	// x, y := "2", 2 // type error
	x, y := 2, "2" // ok
	Println(x, y)
}

func test3() {
	b, err_b := strconv.ParseBool("True")
	f, err_f := strconv.ParseFloat("-1", 64)
	i, err_i := strconv.ParseInt("-1", 10, 64)
	u, err_u := strconv.ParseUint("-1", 10, 64)
	Println(b, err_b)
	Println(f, err_f)
	Println(i, err_i)
	Println(u, err_u)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
