package main

func WaitMe(c chan bool) {
  println("wait!")
  c <- true
}

func main() {
  c := make(chan bool)
  go WaitMe(c)
  <-c
}
