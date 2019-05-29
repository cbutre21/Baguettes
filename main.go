package main

import "fmt"

var Name = "Max"

func F1()string{
  return ("Hello " + F2())
}

func F2()string{
  return Name
}

func main() {
	fmt.Println(F1())
}

