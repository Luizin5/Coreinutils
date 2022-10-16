package main

import ("os/user"
	"log")


func main()  {
    u, err := user.Current()
    if err != nil {
	log.Fatal(err)
    }
    println(u.Username)


}
