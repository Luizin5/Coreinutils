package main

import (
  "flag" 
//  "fmt"
)

func main() {
  help := flag.Bool("help",false,"help message")
  version := flag.Bool("version",false,"print version")
  flag.Parse()

  if *version {
    println(`Coreinutils v0.1.0
    Written by Luizin5`)
  } else if *help {
    println(help_message())
  } else {
    println("Try 'coreinutils -help' for more information.")
  }
}

func help_message() string {
  return `usage of Coreinutils: 
	-help		display this message.
	-version	print version and exit.`
}
