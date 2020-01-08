package main

// go build -ldflags --help
//      -s    disable symbol table
//      -w    disable DWARF generation
//      -v    print link trace
//      -X    definition add string value definition of the form importpath.name=value

// strip size of the exe
// go build -ldflags -s -ldflags -w
// go build -ldflags "-s -w"

// set executable name
// go build -o prog

// without console window
// go build -ldflags "-H=windowsgui"

import "os"
import "fmt"
import "strconv"

// go build -ldflags "-X main.param=abc"
// this will set uninitialized variable to passed value
// this will also set initialized variable to passed value
// var param string
var param string = "foo"

type S0 struct {
	v int
}

func (s S0) show() string {
	return strconv.Itoa(s.v)
}

type S1 struct {
	v int
	S0
}

type S2 struct {
	v int
	S0
}

func (s S2) show() string {
	return "S0.v=" + strconv.Itoa(s.S0.v) + ",S2.v=" + strconv.Itoa(s.v)
}

func test1() {
	fmt.Println("test1")

	s0 := S0{}
	fmt.Println(s0.show())

	s1 := S1{v: 1, S0: S0{v: 2}}
	fmt.Println(s1.show())

	s2 := S2{v: 1, S0: S0{v: 2}}
	fmt.Println(s2.show())
	fmt.Println(s2.S0.show())
}

func test2() {
	fmt.Println("test2")

	S0 := S0{v: 0} // OK
	// Sx := S0{v: 0} // error, S0 is not a type
	// fmt.Println(S0.show() + Sx.show())
	fmt.Println(S0.show())
}

func main() {
	fmt.Printf("param = %s\n", param)
	test1()
	test2()
	os.Exit(1)
}
