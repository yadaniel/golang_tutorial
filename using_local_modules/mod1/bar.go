package mod1

import _ "fmt"
import . "fmt"
import f "fmt"

// import mod "./mod1/mod2"
import mod "mod2"

func init() {
	Printf("mod1::bar.go init")
	f.Printf("mod1::bar.go init")
	println("mod1::bar.go init")
}

func Use_mod2() {
	Printf("%v\n", mod.Y)
}
