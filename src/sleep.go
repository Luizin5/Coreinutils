package main

import (
  "os"
  "time"
  "strconv"
)

func main()  {
    wait_time,err  := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    time.Sleep(time.Duration(wait_time) * time.Second)
}
