package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom-uchida/go-protobuf/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Tomomasa",
		Email:       "tomomasa.is.0930@gmail.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Tomomasa.",
		},
		Birthday: &pb.Date{
			Year:  1995,
			Month: 9,
			Day:   30,
		},
	}

	serializeAndDeserialize(employee)
	jsonMapping(employee)
}

func serializeAndDeserialize(employee *pb.Employee) {
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("failed to serialize", err)
	}

	if err := os.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("failed to write", err)
	}
	in, err := os.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("failed to read", err)
	}

	readEmployee := &pb.Employee{}
	if err := proto.Unmarshal(in, readEmployee); err != nil {
		log.Fatalln("failed to deserialize", err)
	}

	fmt.Println(readEmployee)
}

func jsonMapping(employee *pb.Employee) {
	b, err := protojson.Marshal(employee)
	if err != nil {
		log.Fatalln("failed to marshal to json", err)
	}

	readEmployee := &pb.Employee{}
	if err := protojson.Unmarshal(b, readEmployee); err != nil {
		log.Fatalln("failed to unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
