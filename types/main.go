package main

// golang types and their usage

import "fmt"

func int_types() {
	// all types are defaul initialized with 0
	var u8 uint8
	var _u8 byte // byte is alias for uint8 (aliases have non strict semantic)
	var u16 uint16
	var u32 uint32
	var u64 uint64
	var u uint

	var i8 int8
	var i16 int16
	var i32 int32
	var _i32 rune // rune is alias for int32
	var i64 int64
	var i int

	fmt.Printf("%d,%d,%d,%d,%d,%d\n", u8, _u8, u16, u32, u64, u)
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", i8, i16, i32, _i32, i64, i)
	if u8 == _u8 {
		fmt.Printf("equal\n")
	} else {
		fmt.Printf("not equal\n")
	}
	/*
		// type mismatch compile error
		if u8 == u16 {
			fmt.Printf("equal")
		}
	*/

	// creating new type (strict semantic)
	type ID uint32
	var id ID = 0x00
	fmt.Printf("%d\n", id)

	type Counter uint32
	var cnt Counter = 0x0000
	fmt.Printf("%d\n", cnt)
	/*
		if id == cnt {
			fmt.Printf("does not compile\n", id)
		}
	*/

	/*
		if id == u32 {
			fmt.Printf("does not compile")
		}
	*/
	if id == 0x00 && cnt == 0x00 {
		fmt.Printf("this compiles because literals type depends on variable type\n")
		fmt.Printf("the same is true on variable initialization, when using var <varname> <vartype> = <literal> syntax\n")
	}

	bitmask := 0xFF00 // what type is used when concise variable declaration/definition is used
	fmt.Printf("%d, %b, %x, 0x%X, o%o, %v, %T\n", bitmask, bitmask, bitmask, bitmask, bitmask, bitmask, bitmask)
	fmt.Println("note that all integer types can be formatted with %d,%b,%x,%X,%o and %v ... v accepts any type ")
	fmt.Println("%T gives the variable type") // the type is int
	fmt.Printf("use double %% when using fmt.Printf instead of fmt.Println\n")

	//bitmask = i32	// int is not int32
	bitmask = i

	// let's print the types
	fmt.Printf("%T,%T,%T,%T,%T,%T\n", u8, _u8, u16, u32, u64, u)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n", i8, i16, i32, _i32, i64, i)
}

func bool_type() {
	var b0 bool // default initialized with false
	b1 := true
	fmt.Printf("%t %T\n", b0, b0)
	fmt.Printf("%t %T\n", b1, b1)
}

func float_types() {
	f := 1.123 // default is float64
	fmt.Printf("%f %e %g\n", f, f, f)

	// explicit
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	fmt.Printf("%T,%T,%T,%T\n", f32, f64, c64, c128)
	fmt.Printf("%T,%T,%T,%T\n", real(c64), imag(c64), real(c128), imag(c128))
}

func struct_types() {
	// unnamed struct
	s1 := struct{ x, y, z int32 }{0, 0, 0}
	s2 := struct{ x, y, z int32 }{x: 0, y: 0, z: 0}
	fmt.Printf("%v %+v %T\n", s1, s1, s1) // %+v will print the fields
	fmt.Printf("%v %+v %T\n", s2, s2, s2)
	if s1 == s2 {
		fmt.Println("equal ... unnamed structs are not strict")
	}

	//
	type SA struct {
		x, y, z int32
	}
	type SB = struct { // alternative syntax
		x, y, z int32
	}
	var sa SA // default initialized with 0
	var sb SB
	if sa == sb {
		fmt.Println("equal ... even named structs are not strict")
	}
	if sa == s1 {
		fmt.Println("equal ... unnamed and named structs are not strict")
	}

	// nested structs
	type NS struct {
		vec   struct{ x, y, z uint32 }
		valid bool
	}
	var ns NS
	fmt.Printf("<\n")
	fmt.Printf("%v\n", ns)  //{{0 0 0} false}
	fmt.Printf("%+v\n", ns) //{vec:{x:0 y:0 z:0} valid:false}
	fmt.Printf("%#v\n", ns) //main.NS{vec:struct { x uint32; y uint32; z uint32 }{x:0x0, y:0x0, z:0x0}, valid:false}
	fmt.Printf("%T\n", ns)  //main.NS
	fmt.Printf(">\n")
}

func pointer_types() {
	var pu32 *uint32 // points to 0
	fmt.Printf("%p\n", pu32)
	//fmt.Printf("%d\n", *pu32) // no compile error ... runtime error (panic)
	var u32 uint32 = 100
	pu32 = &u32
	fmt.Printf("%p ... %d = %d\n", pu32, *pu32, u32)
	*pu32 = 200
	fmt.Printf("%p ... %d = %d\n", pu32, *pu32, u32)
}

func string_types() {}

func function_types() {
	var f0 func()
	var f1 func(int) int
	// f0()	// no compile error ... runtime error (panic)
	f0 = func() { fmt.Println("f0 called") }
	f0()
	f1 = func(val int) int { fmt.Println("f1 called"); return val * 2 }
	f1(10) // return value may be ignored
	fmt.Printf("f1(100) = %d\n", f1(100))
	// unnamed function to swap arguments
	p, q := func(x, y int) (a, b int) {
		a = y
		b = x
		return
	}(1, 2)
	fmt.Printf("%d,%d\n", p, q)
	// different equalivalent syntax
	swap := func(x int, y int) (a int, b int) {
		a = y
		b = x
		return
	}
	p, q = swap(3, 4)
	fmt.Printf("%d,%d\n", p, q)
	//pq := swap(3, 4)	// compile error ... multiple value context in single value context ... golang has no tuples
}

func misc() {
	fmt.Printf("\n%d<%d>\n", 1) // 2 arguments expected, 1 given ... this is not compile error ... at runtime gives <%!d(MISSING)>
	fmt.Printf("%d |", 1, 2)    // 1 argument expected, 2 given ... this is not compile error ... at runtime gives |%!(EXTRA int=2)
	fmt.Println()
}

func array_types() {
	// explicit size
	var arr1 [10]byte = [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for idx, val := range arr1 {
		fmt.Printf("arr[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// implicit size
	var arr2 [10]byte = [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var arr2 [10]byte = [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}	// compile error ... size must match
	for idx, val := range arr2 {
		fmt.Printf("arr[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// implicit size and variable declaration ... variable gets the type from [...]byte{} expression
	arr3 := [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array size can be omitted only here
	for idx, val := range arr3 {
		fmt.Printf("arr[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// how are arrays passed to functions
	func(arr [10]byte) { // array size can not be omitted
		arr[0] = 100 // this is local copy, arr3 is not changed
	}(arr3)
	for idx, val := range arr3 {
		fmt.Printf("arr[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// how are arrays passed to functions
	func(arr *[10]byte) {
		arr[0] = 100    // this is reference, arr3 is changed ... syntax 1
		(*arr)[0] = 200 // this is reference, arr3 is changed ... syntax 2
	}(&arr3)
	for idx, val := range arr3 {
		fmt.Printf("arr[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// builtin function on array ... len (length)
	fmt.Printf("len(arr3)=%d\n", len(arr3))
	// arrays can be compared when the arrays have the same type ... [N]type
	if arr1 == arr2 {
		fmt.Println("[arr1==arr2] equal")
	}
	arr4 := [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	fmt.Printf("arr4 = %v\n", arr4)
	/*
		// compile error ... mismatched types [10]byte and [15]byte
		if arr1 == arr4 {
			fmt.Println("[arr1==arr4] equal")
		} else {
			fmt.Println("[arr1!=arr4] not equal")
		}
	*/
	// array assignment semantic (value semantic)
	arr5 := [...]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr6 := arr5 // copy is created
	fmt.Printf("arr5 = %v ... arr6 = %v\n", arr5, arr6)
	arr6[0] = 10 // arr5 not changed
	fmt.Printf("arr5 = %v ... arr6 = %v\n", arr5, arr6)
	// array can not be assigned nil
	//arr6 = nil
	// empty array ... the set of the array type have special value, empty array
	arr7 := [0]uint32{}
	arr8 := [...]uint32{}
	if arr7 == arr8 {
		fmt.Println("[arr7 == arr8] empty array are equal")
	}
	fmt.Printf("[arr7,arr8] have types %T ... %T\n", arr7, arr8)
}

func slice_types() {
	// non initialized slice
	var s0 []int // is empty and safe to use
	for idx, val := range s0 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// slice with 0 elements ... initialized with zero len slice
	var s1 []int = []int{} // is empty and safe to use
	for idx, val := range s1 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// slice with 0 elements ... initialized with builtin make function
	var s2 []int = make([]int, 0) // is empty and safe to use
	for idx, val := range s2 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// slice with 10 elements ... initialized with slice
	var s3 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for idx, val := range s3 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// slice with 10 elements ... initialized with make function
	var s4 []int = make([]int, 10) // all elements are default initialized with 0
	for idx, val := range s4 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// builtin functions on slices ... len (length), cap (capacity), append
	s5 := make([]int, 10)
	fmt.Printf("len(s5)=%d, cap(s5)=%d\n", len(s5), cap(s5))
	//
	s5 = make([]int, 10, 20)
	fmt.Printf("len(s5)=%d, cap(s5)=%d\n", len(s5), cap(s5))
	//
	s5 = append(s5, []int{10, 11, 12, 13, 14}...)
	fmt.Printf("len(s5)=%d, cap(s5)=%d\n", len(s5), cap(s5))
	for idx, val := range s5 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// append above capacity ... capacity is doubled when exceeded
	s5 = append(s5, s5...)
	fmt.Printf("len(s5)=%d, cap(s5)=%d\n", len(s5), cap(s5))
	for idx, val := range s5 {
		fmt.Printf("slice[%d]=%d, ", idx, val)
	}
	fmt.Println()
	// slice indexing
	s6 := make([]uint32, 10, 20) // len = 10, cap = 20
	// from 0 to stop
	s7 := s6[0:5] // from 0 to 5 (excluding), cap remains 20
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// from default 0 to stop
	s7 = s6[:5] // from 0 to 5 (excluding), cap remains 20
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// from start to stop
	s7 = s6[5:10] // from 5 to 10 (excluding), cap now 15
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// from start to default stop (len()+1)
	s7 = s6[5:] // from 5 to 10 (excluding), cap now 15
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// from start to stop
	s7 = s6[2:8] // cap now 18
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// image slice as data structure comprised of 3 pointers ... pstart, pstop (pointing to valid data) and pmax (pointing to allocated uninitialized memory)
	// len being pstop-pstart+1, cap being pmax-pstart+1 (allocation logic not described)
	// when subslice is taken with syntax slice[a:b], pstart of the sublice becomes pstart+a ... pstop, pmax remain ... len, cap change accordingly
	s7 = s6[2:8:8] // 7 elements, cap now 6 ... cap is decreased
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// when subslice is taken with syntax slice[a:b:c] the possibility is given to cut the cap
	// pstart of the sublice becomes pstart+a, pstop becomes pstop+b, pmax becomes c-a
	s7 = s6[2:8:10] // 7 elements, cap now 8 ... cap is decreased
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	//
	//s7 = s6[2:8:30] // 7 elements, cap now 28 .. cap can not be increased ... this is runtime error (panic)
	//fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	s7 = s6[2:8:20] // 7 elements, cap 18 as it would be when c is omitted
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	s7 = s6[2:8] // 7 elements, cap 18
	fmt.Printf("len()=%d, cap()=%d\n", len(s7), cap(s7))
	// taking slice from array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%T ... %v\n", arr, arr)
	s8 := arr[0:4] // the result is slice, cap = len(arr)
	fmt.Printf("%T ... %v, len()=%d, cap()=%d\n", s8, s8, len(s8), cap(s8))
	s8 = arr[0:4:8] // all slice indices can be used, cap = len(c-a) ... here 8 which is len(arr)
	fmt.Printf("%T ... %v, len()=%d, cap()=%d\n", s8, s8, len(s8), cap(s8))
	s8 = arr[0:4:5] // all slice indices can be used, cap = len(c-a) ... here 5
	fmt.Printf("%T ... %v, len()=%d, cap()=%d\n", s8, s8, len(s8), cap(s8))
	// slice assignment semantics (reference semantic)
	s9 := []uint32{0, 1, 2, 3, 4, 5, 7, 8, 9}
	s10 := s9 // s10 is reference, not copy
	/*
		// 2 slices can not be compared
		if s9 == s10 {
			fmt.Println("slices equal")
		}
	*/
	fmt.Printf("[s9] %v ... %v\n", s9, s10)
	s10[0] = 10 // change s10, both change
	fmt.Printf("[s9] %v ... %v\n", s9, s10)
	//
	s11 := []uint32{} // slice is empty but not nil
	if s11 == nil {
		fmt.Println("[s11] slice is nil")
	} else {
		fmt.Println("[s11] slice is not nil")
	}
	s11 = nil // slice variable can be assigned nil
	if s11 == nil {
		fmt.Println("[s11 nil] slice is nil")
	} else {
		fmt.Println("[s11 nil] slice is not nil")
	}
	// slice type does not include the len and cap, only the []type
	s12 := []uint32{}
	s13 := []uint32{0}
	s14 := make([]uint32, 0, 10)
	fmt.Printf("[s12,s13,s14] %T,%T,%T\n", s12, s13, s14)
}

func map_types() {
	var m1 map[int]string       // default is nil
	var m2 map[int]string = nil // explicit initialization
	fmt.Printf("[m1,m2] %T ... %T\n", m1, m2)
	if m1 == nil {
		fmt.Println("map is inizialized to nil ... default")
	}
	/*
		// map can only be compared to nil
		if m1 == m2 {
			fmt.Println("map can not be compared")
		}
	*/
	for key, val := range m2 {
		fmt.Printf("[map] %v -> %v\n", key, val)
	}
}

func variable_declaration() {
	// var syntax
	var x1 int                                                // default initialized with 0
	var x2 int = 10                                           // initialized with 10
	var x3, x4 int                                            // both default initialized with 0
	var x5, x6 int = 100, 200                                 // both initialized ... 100 and 200
	_, _, _, _, _, _ = x1, x2, x3, x4, x5, x6                 // discard the value, compiler requires that variables are used
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", x1, x2, x3, x4, x5, x6) // or use them
	// var block syntax
	var (
		x7       int
		x8       int = 10
		x9, x10  int
		x11, x12 int = 100, 200
	)
	_, _, _, _, _, _ = x7, x8, x9, x10, x11, x12
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", x7, x8, x9, x10, x11, x12)
}

func const_declaration() {
	// const syntax
	//const X0	// initialization can not be omitted
	const X1 = 0       // implicit type = int (default for numerical non-floating point literal)
	const X2 int32 = 0 // explicit type, literal is casted to int32
	const X3, X4 = 1, 2
	const X5, X6 uint32 = 1, 2
	//note: compiler does not require usage (read access not required)
	// printed only to confirm the values
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", X1, X2, X3, X4, X5, X6)
	// const block syntax
	const (
		X7              = 0
		X8       uint32 = 0
		X9, X10         = 1, 2
		X11, X12 uint32 = 1, 2
	)
	fmt.Printf("%d,%d,%d,%d,%d,%d\n", X7, X8, X9, X10, X11, X12)
}

func type_alias() {
	type ID uint32     // strict type
	type GUID = uint32 // alias type
	var u32 uint32 = 10
	var id ID = 10
	var guid GUID = 10
	/*
		if u32 == id {	// compile error ... type mismatch uint32 and ID
			fmt.Println("equal")
		}
	*/
	if u32 == guid {
		fmt.Println("equal")
	}
	fmt.Printf("%d,%d,%d\n", id, guid, u32)
}

func main() {
	int_types()
	bool_type()
	float_types()
	string_types()
	struct_types()
	pointer_types()
	function_types()
	array_types()
	slice_types()
	map_types()
	variable_declaration()
	const_declaration()
	type_alias()
	misc()
}
