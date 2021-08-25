package main

import "fmt"

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("woof")
}

type Husky struct {
	Dog
}

func main() {
	h := Husky{Dog{}}
	h.Speak()
}
