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
}

func array_types() {}
func slice_types() {}
func map_types()   {}

func main() {
	int_types()
	bool_type()
	float_types()
	struct_types()
	pointer_types()
	function_types()
	array_types()
	slice_types()
	map_types()
	misc()
}
