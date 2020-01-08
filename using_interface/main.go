package main

import "fmt"

type I0 interface {
	use_v()
}

type I1 interface {
	use_v()
	// f()
}

type S0 struct {
	v int
}

// this function implements both interfaces I0 and I1 at once
// when f() remains commented
func (s S0) use_v() {
	fmt.Println("S0.use_v()")
}

type S1 struct {
	v int
}

// this function implements both interfaces I0 and I1 at once
// when f() remains commented
func (s S1) use_v() {
	fmt.Println("S1.use_v()")
}

func takes_iface0(i I0) {
	i.use_v()
}

func takes_iface1(i I1) {
	i.use_v()
}

// type must implement all functions of the interface
func main() {
	s0 := S0{v: 100}
	s1 := S1{v: 200}
	takes_iface0(s0)
	takes_iface0(s1)
	takes_iface1(s0) // error, when f() is uncommented
	takes_iface1(s1) // error, when f() is uncommented
}
