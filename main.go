package main

import (
	"fmt"
	"go-protocol-buffer/src/person/personpb"
)

func main() {
	doSimple()
	fmt.Print("hello")
}

func doSimple() {
	sm := &personpb.Person{
		Name:  "serin",
		Id:    2,
		Email: "dd",
	}
	fmt.Print(sm)
}
