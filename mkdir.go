package main

import "os"
import "log"

func main()  {
	var err error
	for _,i := range os.Args[1:]  {
		err = os.Mkdir(i,0750)
		if err != nil {
			log.Fatal(err)
		}
	}
}
