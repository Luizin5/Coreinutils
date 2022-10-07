package main

import "os"

func main()  {
	for _,i:= range os.Args  {
		os.Create(i)
	}
}
