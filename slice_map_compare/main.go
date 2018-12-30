package main

import "fmt"

func slice_map_properties_compare() {
	// equal properties -> declaration and definition
	// slice
	var s0 []uint32
	if s0 == nil {
		fmt.Println("default nil")
	}
	s1 := []uint32{}
	if s1 != nil {
		fmt.Println("empty, but not nil")
	}
	s2 := make([]uint32, 0)
	if s2 != nil {
		fmt.Println("empty, but not nil")
	}
	// map
	var m0 map[int]uint32
	if m0 == nil {
		fmt.Println("default nil")
	}
	m1 := map[int]uint32{}
	if m1 != nil {
		fmt.Println("empty, but not nil")
	}
	m2 := make(map[int]uint32, 0)
	if m2 != nil {
		fmt.Println("empty, but not nil")
	}
	fmt.Println()
	// equal properties -> reference semantic
	// slice
	s3 := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := s3
	fmt.Printf("%v = %v\n", s3, s4)
	fmt.Printf("%p != %p, %v = %v\n", &s3, &s4, &s3, &s4) // pointers are not equal!!!
	s3[0] = 10                                            // changes s4 as well
	s4[9] = 0                                             // changes s3 as well
	fmt.Printf("%v = %v\n", s3, s4)
	fmt.Printf("%p != %p, %v = %v\n", &s3, &s4, &s3, &s4)
	// map
	m3 := map[int]uint32{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	m4 := m3
	fmt.Printf("%v = %v\n", m3, m4)
	fmt.Printf("%p != %p, %v = %v\n", &m3, &m4, &m3, &m4) // pointers are not equal!!!
	m3[0] = 10                                            // changes m4 as well
	m4[9] = 0                                             // changes m3 as well
	fmt.Printf("%v = %v\n", m3, m4)
	fmt.Printf("%p != %p, %v = %v\n", &m3, &m4, &m3, &m4)
	fmt.Println()
	// equal properties -> can not be compared to other object of the same type (including itself) ... compare to nil ok
	// slice
	//
	//if []uint32{1, 2, 3} == []uint32{1, 2, 3} {
	//}
	// map
	//if map[int]uint32{0: 0} == map[int]uint32{0: 0} {
	//}
	var snil []uint32
	var mnil map[int]uint32
	var sb1 bool = snil == nil
	var mb1 bool = mnil == nil
	var sb2 bool = []uint32{} == nil
	var mb2 bool = map[int]uint32{} == nil
	fmt.Printf("%v,%v,%v,%v\n", sb1, mb1, sb2, mb2) // true, true, false, false
	fmt.Println()
	// unequal properties -> take address from container
	// slice
	s := []uint32{1, 2, 3}
	ptr_into_slice := &s[0]            // can take address from slice element ... not recommended
	fmt.Printf("%v\n", ptr_into_slice) // when slice is relocated, the old memory can not be freed
	// map
	//m := map[int]uint32{1: 1, 2: 2, 3: 3}
	//ptr_into_map := &m[1] // can not take address from map element
	//fmt.Printf("%v\n", ptr_into_map)
	// unequal properties -> assignment to struct in container
	type S0 struct {
		x0, x1, x2 uint32
	}
	// struct
	st := S0{}
	st.x0 = 100
	// slice of structs
	sl := []S0{S0{}}
	sl[0].x0 = 100
	// map
	mp := map[int32]S0{0: S0{}}
	//mp[0].x0 = 100    // can not assign to struct field in map
	mp[0] = S0{x0: mp[0].x0, x1: mp[0].x1, x2: mp[0].x2} // workaround
}

func main() {
	slice_map_properties_compare()
}
