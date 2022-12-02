package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
	    for { println("y") }
	} else {
	    for { 
	      println( strings.Join( args[1:]," ") ) 
	    }
	}
}
