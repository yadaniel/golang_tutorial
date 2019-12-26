module main

go 1.13

require (
	mod1 v0.0.0
	mod2 v0.0.0
)

replace (
	mod1 v0.0.0 => ../mod1
	mod2 v0.0.0 => ../mod2
)
