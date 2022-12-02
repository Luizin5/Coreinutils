package main

import (
"os"
"flag"
//"log"
)

func main()  {
      
	directory := os.Args[2]
	option := flag.Bool("rf",false,"remove directories that's not empty")
	flag.Parse()
	
	if !del(directory,option)  {
		println("error")
	}
}

func del(dir string, option * bool) bool {
    fileinfo, _ := os.Stat(dir)

    if *option && fileinfo.IsDir() {
	os.RemoveAll(dir)
	return true
    } else if !fileinfo.IsDir() {
	os.Remove(dir)
	return true
    }
    return false
}

