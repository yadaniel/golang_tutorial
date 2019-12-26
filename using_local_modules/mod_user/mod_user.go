package main

import "mod1"
import "mod2"

var x mod2.X

func main() {
    mod1.Use_mod2()
}
