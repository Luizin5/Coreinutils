package main

import (
	Args "coreutils/getargs"
)

func main() {
	var args = Args.Get()
	if args == "" {
		print_loop("y")
	} else {
		print_loop(args)
	}
}

func print_loop(str string)  {
	for { println(str) }
}
