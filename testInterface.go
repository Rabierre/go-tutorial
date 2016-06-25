package main

type Soundable interface {
    MakeNoise() string
}

type Duck struct {}

func (d Duck) MakeNoise() string {
    return "Quak"
}

type Ship struct {}

func (s Ship) MakeNoise() string {
    return "Bahhhhh"
}

func main() {
  println(Duck{}.MakeNoise())
  println(Ship{}.MakeNoise())
}
