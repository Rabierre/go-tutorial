package main

func Greet(c chan string) {
    c <- "Hello Goroutine!"
}

func main() {
  c := make(chan string)
  go Greet(c)
  println(<-c)
}
