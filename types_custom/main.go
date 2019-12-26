package main

import "os"

// types defined with = are synonyms
// structs with same structure can be used instead

type I1 interface{ foo() int }
type I2 = interface{ foo() int }

// S0 and S1 are different types and can not be mixed
// S2, S3 can be used where S0 or S1 are expected
// S0, S1 can be used where S2 or S2 are expected
type S0 struct{ x int }
type S1 struct{ x int }
type S2 = struct{ x int }
type S3 = struct{ x int }

var s00 S0 = S0{}

/*  var s01 S0 = S1{} */
var s02 S0 = S2{}
var s03 S0 = S3{}

// var s10 S1 = S0{}
var s11 S1 = S1{}
var s12 S1 = S2{}
var s13 S1 = S3{}

var s20 S2 = S0{}
var s21 S2 = S1{}
var s22 S2 = S2{}
var s23 S2 = S3{}

var s30 S3 = S0{}
var s31 S3 = S1{}
var s32 S3 = S2{}
var s33 S3 = S3{}

// here types defined with = are not considered aliases to types defined without =
// types defined with = are considered aliases to types defined with = only
// this is different compared to structs
type Id0 int8
type Id1 int8
type Id2 = int8
type Id3 = int8

func compare(id1 Id1, id2 Id2) bool {
	return true
}

func compare2(id2 Id2, id3 Id3) bool {
	return true
}

func takes_S1(s S1)       {}
func takes_S1_S2_S3(s S2) {}

func takes_I1(i I1) {}
func takes_I2(i I2) {}

func (s *S1) foo() int {
	return (*s).x
}

// func (s *S1) foo() (x int) {
// 	x = (*s).x
// 	return
// }

// type S2 = struct{ x int }
// S2 is not defined type ... just struct {x int}
// func (s *S2) foo() (x int) {
// 	x = 1
// 	return
// }

func main() {
	id1 := Id1(0)
	id2 := Id2(0)
	compare(id1, id2)

	// this error is catched
	// compare(id2, id1)

	id3 := Id3(0)
	compare2(id2, id3)

	// no difference
	takes_S1(S1{})
	takes_S1(S2{})
	takes_S1(S3{})
	takes_S1_S2_S3(S1{})
	takes_S1_S2_S3(S2{})
	takes_S1_S2_S3(S2{})

	s1 := S1{}
	s2 := S2{}
	s3 := S3{}
	takes_S1(s1)
	takes_S1(s2)
	takes_S1(s3)
	takes_S1_S2_S3(s1)
	takes_S1_S2_S3(s2)
	takes_S1_S2_S3(s3)

	// takes_I1(&s1)
	// takes_I2(&s2)

	os.Exit(0)
}
