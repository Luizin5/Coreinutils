package main

import (
"os"
"io/ioutil"
)


func main()  {
	args := os.Args
	switch len(args) {
	  case 1:
	    show(true,args)
	  case 2:
	    show(false,args)
	  default:
	    println("invalid amout of arguments")
	}
}

func show(this_dir bool, args []string)  {
	if this_dir {
		files,err := ioutil.ReadDir("./")
		switch err {
		  case nil:
			for _,i := range files {
				println(i.Name())
			}
		  default:
			println(err)
		}
	} else {
		files,err := ioutil.ReadDir(args[1])
		switch err {
		  case nil:
			  for _,i := range files {
				println(i.Name())
			  }
		  default:
			  println(err)
		}
	}
}
