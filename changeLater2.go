package main

type Number struct {
  Item int
}

func UpdateLater(c chan Number) {
  number := Number{Item: 1}
  c <- number

  number.Item = 2
  println("In UpdateLater Function: ", number.Item)
}

func main() {
  c := make(chan Number)
  go UpdateLater(c)

  result := <-c
  println("Main Function: ", result.Item)
}
