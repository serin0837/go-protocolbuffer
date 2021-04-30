package main

import (
	"fmt"
	"go-protocol-buffer/src/simple/simplepb"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "my name",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	sm.Name = "I rename you"
	fmt.Println(sm)
	fmt.Println("The Id is:", sm.GetId())
}
