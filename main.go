package main

import "fmt"

var Name1 = "Max"
var Name2 = "Celian"

func F1()string{
  return ("Hello " + F2())
}

func F2()string{
  return (Name1 + " I am " + F3())
}

func F3()string{
  return Name2
}

func main() {
	fmt.Println(F1())
}
