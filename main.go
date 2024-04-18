package main

import (
	"fmt"
	"log"
	"os"

	"github.com/okaberintaroubeta/proto_example/src/simple/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()
	fname := "test.bin"
	writeToFile(fname, sm)
	sm2 := &simplepb.SimpleMessage{}
	err := readFromFile(fname, sm2)
	if err != nil {
		log.Fatalln("Error reading file")
	}
	jsonDemo()

}

func jsonDemo() {
	m := doSimple()
	jsonStr := toJson(m)
	fmt.Println("To JSON:", jsonStr)
	m2 := &simplepb.SimpleMessage{}
	fromJson([]byte(jsonStr), m2)
	fmt.Println("From JSON:", m2)
}

func readFromFile(fname string, m proto.Message) error {
	output, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading from file", fname)
		return err
	}
	err = proto.Unmarshal(output, m)
	if err != nil {
		log.Fatalln("Error reading message", err)
		return err
	}
	return nil
}

func writeToFile(fname string, m proto.Message) error {
	content, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(fname, content, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	return nil
}

func toJson(m proto.Message) string {
	content, err := protojson.Marshal(m)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return string(content)
}

func fromJson(content []byte, m proto.Message) error {
	err := protojson.Unmarshal(content, m)
	if err != nil {
		log.Fatalln("Can't convert from JSON", err)
		return err
	}
	return nil

}

func doSimple() *simplepb.SimpleMessage {
	sm := &simplepb.SimpleMessage{
		Id:          12345,
		IsSimple:    true,
		Name:        "My simple mesage",
		SimpleLists: []int32{1, 4, 5, 7},
	}
	return sm
}
