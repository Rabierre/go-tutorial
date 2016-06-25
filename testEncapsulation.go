package main

import (
    "./spy"
)

func main() {
    s := spy.Spy{}
    s.PublicName = "Angel"
    s.privateName = "Evil"
}
