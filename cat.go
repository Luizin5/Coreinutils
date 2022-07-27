package main

import (
	"os"
)

func main()  {
	if len(os.Args) != 2 { 
		println("invalid amount of arguments")
		os.Exit(0)
	}
	content,err := os.ReadFile(os.Args[1])
	if err == nil {
		println(string(content))
	} else {
		println(err)
	}
}
