package main

import "os"
import "path/filepath"

func main()  {
	file,err := os.Executable()
	if err != nil {
		panic(err)
	}
	println(filepath.Dir(file))
}
