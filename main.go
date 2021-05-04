package main

import (
	"fmt"
	"go-protocol-buffer/src/simple/simplepb"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
}

//did a bit differently with tutorial, directrly pointing to struct of protocol buffer
//json for debugging for..
func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	sm := &simplepb.SimpleMessage{
		Id:         123456,
		IsSimple:   true,
		Name:       "my name",
		SampleList: []int32{1, 4, 7, 8},
	}

	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(sm)
	if err != nil {
		log.Fatalln("can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, sm *simplepb.SimpleMessage) {
	err := jsonpb.UnmarshalString(in, sm)
	if err != nil {
		log.Fatalln("couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	//file mode golang
	//ioutil
	//A FileMode represents a file's mode and permission bits.
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("data has been written")
	return nil
}
func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something went wrong", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Coudnt put the bytes into the protocol buffers")
		return err2
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
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

	return &sm
}

//Package proto provides functionality for handling protocol buffer messages. In particular, it provides marshaling and unmarshaling between a protobuf message and the binary wire format.
