package main 

import "time"

func main() {
    go println("Hello Goroutine")
    time.Sleep(time.Second)
}
