package main

func UpdateLater(c chan int) {
  number := 1
  c <- number

  number = 2
}

func main() {
  c := make(chan int)
  go UpdateLater(c)
  println(<-c)
}
