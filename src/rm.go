package main

import (
"os"

//"log"
)

func main()  {
	args := os.Args
	var directory string
	var option string
	if args[1] == "-rf"{
		directory = args[2]
		option = args[1]
	} else {
		directory = args[1]
		option = ""
	}
	if !del(directory,option)  {
		println("error")
	}
}

func del(dir, option string) bool {
    fileinfo, _ := os.Stat(dir)

    if option == "-rf" && fileinfo.IsDir() {
	os.RemoveAll(dir)
	return true
    } else if option != "-rf" && !fileinfo.IsDir()  {
	os.Remove(dir)
	return true
    }
    return false
}

